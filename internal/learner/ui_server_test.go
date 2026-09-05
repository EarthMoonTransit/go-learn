package learner

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"testing"
	"time"

	"go-practice-bank/internal/tasking"
)

type uiServerOutput struct {
	started  chan string
	stopping chan struct{}
}

func (o *uiServerOutput) Write(data []byte) (int, error) {
	text := string(data)
	if strings.HasPrefix(text, "Go Learn UI: ") {
		o.started <- strings.Fields(text)[3]
	}
	if strings.HasPrefix(text, "Stopping UI;") {
		close(o.stopping)
	}
	return len(data), nil
}

func TestServeUIShutdownDrainsVerificationAndPersistsActualResult(t *testing.T) {
	s := uiTestService(t, "")
	if _, err := s.Start("literal-bindings-001"); err != nil {
		t.Fatal(err)
	}
	entered, release := make(chan struct{}), make(chan struct{})
	var once sync.Once
	unblock := func() { once.Do(func() { close(release) }) }
	s.verify = func(_, _, _ string, scope tasking.LearnerVerificationScope) (tasking.LearnerVerificationResult, error) {
		close(entered)
		<-release
		return tasking.LearnerVerificationResult{Scope: scope, Passed: false, Summary: "Fixture full verification failed."}, nil
	}
	ctx, cancel := context.WithCancel(context.Background())
	output := &uiServerOutput{make(chan string, 1), make(chan struct{})}
	done := make(chan error, 1)
	go func() { done <- ServeUI(ctx, s, 0, output) }()
	t.Cleanup(func() { unblock(); cancel() })
	var printed string
	select {
	case printed = <-output.started:
	case err := <-done:
		t.Fatalf("server failed: %v", err)
	case <-time.After(5 * time.Second):
		t.Fatal("server did not start")
	}
	address, err := url.Parse(printed)
	if err != nil {
		t.Fatal(err)
	}
	if address.Hostname() != "127.0.0.1" {
		t.Fatalf("non-loopback bind: %s", address.Hostname())
	}
	token := address.Fragment
	address.Fragment = ""
	origin := strings.TrimSuffix(address.String(), "/")
	client := &http.Client{Timeout: 5 * time.Second}
	body := uiActionBody(t, s, "complete", "literal-bindings-001", "request-shutdown-0001")
	request, err := http.NewRequest("POST", origin+"/api/action", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Origin", origin)
	request.Header.Set("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		t.Fatal(err)
	}
	var admitted uiJob
	err = json.NewDecoder(response.Body).Decode(&admitted)
	response.Body.Close()
	if err != nil || response.StatusCode != 202 || admitted.State != "running" {
		t.Fatalf("admission: %+v %v", admitted, err)
	}
	select {
	case <-entered:
	case <-time.After(5 * time.Second):
		t.Fatal("verification not started")
	}
	cancel()
	select {
	case <-output.stopping:
	case <-time.After(5 * time.Second):
		t.Fatal("shutdown not started")
	}
	select {
	case err := <-done:
		t.Fatalf("server orphaned active verifier: %v", err)
	case <-time.After(30 * time.Millisecond):
	}
	// Once stopping has begun, HTTP must either reject admission or be closed.
	request, err = http.NewRequest("POST", origin+"/api/action", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Origin", origin)
	request.Header.Set("Content-Type", "application/json")
	if response, err = client.Do(request); err == nil {
		io.Copy(io.Discard, response.Body)
		response.Body.Close()
		if response.StatusCode != 503 {
			t.Fatalf("shutdown admitted operation: %d", response.StatusCode)
		}
	}
	unblock()
	select {
	case err := <-done:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("server did not drain")
	}
	p, err := s.Progress()
	if err != nil {
		t.Fatal(err)
	}
	if task := p.Tasks["literal-bindings-001"]; task.Completed || task.Attempts != 1 || task.LastCompletionSummary != "Fixture full verification failed." {
		t.Fatalf("shutdown lost actual result: %+v", task)
	}
	if err := s.Pause(); err != nil {
		t.Fatalf("shutdown left operation lease: %v", err)
	}
}
