package learner

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"

	"go-practice-bank/internal/tasking"
)

func testLocalUI(t *testing.T, s *Service) *localUI {
	t.Helper()
	u, err := newLocalUI(s, "127.0.0.1:43210")
	if err != nil {
		t.Fatal(err)
	}
	return u
}

func uiHTTP(u *localUI, method, path, body string, mutate func(*http.Request)) *httptest.ResponseRecorder {
	r := httptest.NewRequest(method, "http://"+u.host+path, strings.NewReader(body))
	r.Header.Set("Authorization", "Bearer "+u.token)
	r.Header.Set("Origin", "http://"+u.host)
	r.Header.Set("Content-Type", "application/json")
	if mutate != nil {
		mutate(r)
	}
	w := httptest.NewRecorder()
	u.ServeHTTP(w, r)
	return w
}

func uiActionBody(t *testing.T, s *Service, action, id, requestID string) string {
	t.Helper()
	view, err := s.TaskView(id)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(uiRequest{action, id, view.Revision, requestID})
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func waitUIJob(t *testing.T, u *localUI, id string) uiJob {
	t.Helper()
	u.workers.Wait()
	w := uiHTTP(u, "GET", "/api/job?id="+id, "", nil)
	if w.Code != 200 {
		t.Fatalf("job HTTP %d: %s", w.Code, w.Body)
	}
	var job uiJob
	if err := json.Unmarshal(w.Body.Bytes(), &job); err != nil {
		t.Fatal(err)
	}
	return job
}

func TestUIHTTPRejectsUnauthorizedAndMalformedRequestsWithoutMutation(t *testing.T) {
	s := uiTestService(t, "")
	if _, err := s.Start("literal-bindings-001"); err != nil {
		t.Fatal(err)
	}
	u := testLocalUI(t, s)
	body := uiActionBody(t, s, "hint", "literal-bindings-001", "request-boundary-0001")
	before := persistedBytes(t, s)
	tests := []struct {
		name, method, path, body string
		status                   int
		mutate                   func(*http.Request)
	}{
		{"private state", "GET", "/api/state", "", 401, func(r *http.Request) { r.Header.Del("Authorization") }},
		{"private task", "GET", "/api/task?id=literal-bindings-001", "", 401, func(r *http.Request) { r.Header.Set("Authorization", "Bearer wrong") }},
		{"host", "POST", "/api/action", body, 403, func(r *http.Request) { r.Host = "evil.example" }},
		{"origin", "POST", "/api/action", body, 403, func(r *http.Request) { r.Header.Set("Origin", "https://evil.example") }},
		{"missing origin", "POST", "/api/action", body, 403, func(r *http.Request) { r.Header.Del("Origin") }},
		{"fetch site", "POST", "/api/action", body, 403, func(r *http.Request) { r.Header.Set("Sec-Fetch-Site", "cross-site") }},
		{"GET mutation", "GET", "/api/action", "", 404, nil},
		{"PUT mutation", "PUT", "/api/action", body, 405, nil},
		{"form", "POST", "/api/action", body, 400, func(r *http.Request) { r.Header.Set("Content-Type", "application/x-www-form-urlencoded") }},
		{"trailing JSON", "POST", "/api/action", body + " {}", 400, nil},
		{"oversized body", "POST", "/api/action", strings.Repeat(" ", 4097) + body, 400, nil},
		{"unknown field", "POST", "/api/action", strings.TrimSuffix(body, "}") + `,"workspace":"/tmp/arbitrary"}`, 400, nil},
		{"unknown action", "POST", "/api/action", strings.Replace(body, `"hint"`, `"solution"`, 1), 400, nil},
		{"query action", "POST", "/api/action?root=/tmp", body, 400, nil},
		{"bad task ID", "GET", "/api/task?id=../../solution/reference", "", 400, nil},
		{"duplicate task ID", "GET", "/api/task?id=literal-bindings-001&id=zero-results-001", "", 400, nil},
		{"state query", "GET", "/api/state?home=/tmp", "", 400, nil},
		{"oversized URI", "GET", "/api/task?id=" + strings.Repeat("a", 2100), "", 414, nil},
		{"reference endpoint", "GET", "/api/solution", "", 404, nil},
		{"reference path", "GET", "/tasks/ready/literal-bindings-001/solution/reference/bindings.go", "", 404, nil},
		{"traversal", "GET", "/../../progress.json", "", 404, nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := uiHTTP(u, test.method, test.path, test.body, test.mutate)
			if w.Code != test.status {
				t.Fatalf("HTTP %d want %d: %s", w.Code, test.status, w.Body)
			}
			if w.Header().Get("Cache-Control") != "no-store" {
				t.Fatal("private response is cacheable")
			}
			if strings.Contains(w.Body.String(), s.home) || strings.Contains(w.Body.String(), s.root) {
				t.Fatal("internal path leaked")
			}
		})
	}
	if !bytes.Equal(before, persistedBytes(t, s)) {
		t.Fatal("rejected request mutated progress")
	}
	if len(u.jobs) != 0 {
		t.Fatal("rejected request admitted a job")
	}
}

func TestUIPublicShellContainsNoPrivateState(t *testing.T) {
	s := uiTestService(t, "")
	u := testLocalUI(t, s)
	for _, path := range []string{"/", "/app.js", "/app.css"} {
		w := uiHTTP(u, "GET", path, "", func(r *http.Request) { r.Header.Del("Authorization") })
		if w.Code != 200 {
			t.Fatalf("asset %s: %d", path, w.Code)
		}
		for _, private := range []string{s.home, s.root, u.token, "literal-bindings-001"} {
			if strings.Contains(w.Body.String(), private) {
				t.Fatalf("private state in %s", path)
			}
		}
		if !strings.Contains(w.Header().Get("Content-Security-Policy"), "frame-ancestors 'none'") {
			t.Fatal("missing frame policy")
		}
	}
}

func TestUIJobsDeduplicateCoordinateAndDrain(t *testing.T) {
	s := uiTestService(t, "")
	if _, err := s.Start("literal-bindings-001"); err != nil {
		t.Fatal(err)
	}
	u := testLocalUI(t, s)
	entered, release := make(chan struct{}), make(chan struct{})
	var calls atomic.Int32
	s.verify = func(_, _, _ string, scope tasking.LearnerVerificationScope) (tasking.LearnerVerificationResult, error) {
		calls.Add(1)
		close(entered)
		<-release
		return tasking.LearnerVerificationResult{Passed: true, Scope: scope, Summary: "Fixture pass"}, nil
	}
	defer func() { close(release); u.workers.Wait() }()
	id := "request-running-0001"
	body := uiActionBody(t, s, "check", "literal-bindings-001", id)
	if w := uiHTTP(u, "POST", "/api/action", body, nil); w.Code != 202 {
		t.Fatalf("first admission: %d %s", w.Code, w.Body)
	}
	<-entered
	if w := uiHTTP(u, "POST", "/api/action", body, nil); w.Code != 200 {
		t.Fatalf("duplicate: %d", w.Code)
	}
	if w := uiHTTP(u, "POST", "/api/action", strings.Replace(body, id, "request-running-0002", 1), nil); w.Code != 409 {
		t.Fatalf("second tab: %d", w.Code)
	}
	if w := uiHTTP(u, "POST", "/api/action", strings.Replace(body, `"check"`, `"complete"`, 1), nil); w.Code != 409 {
		t.Fatalf("payload binding: %d", w.Code)
	}
	other := uiTestService(t, s.home)
	if _, err := other.Check(); !errors.Is(err, ErrOperationBusy) {
		t.Fatalf("CLI overlap: %v", err)
	}
	if w := uiHTTP(u, "GET", "/api/state", "", nil); w.Code != 200 || !strings.Contains(w.Body.String(), `"state":"running"`) {
		t.Fatalf("running state missing: %s", w.Body)
	}
	u.beginShutdown()
	if w := uiHTTP(u, "POST", "/api/action", body, nil); w.Code != 503 {
		t.Fatalf("shutdown admission: %d", w.Code)
	}
	if calls.Load() != 1 {
		t.Fatalf("duplicate verification: %d", calls.Load())
	}
}

func TestUIRedactsErrorsAndFullDiagnostics(t *testing.T) {
	for _, internalError := range []bool{false, true} {
		t.Run(fmt.Sprint(internalError), func(t *testing.T) {
			s := uiTestService(t, "")
			if _, err := s.Start("literal-bindings-001"); err != nil {
				t.Fatal(err)
			}
			s.verify = func(_, _, _ string, scope tasking.LearnerVerificationScope) (tasking.LearnerVerificationResult, error) {
				if internalError {
					return tasking.LearnerVerificationResult{}, errors.New("PRIVATE_GRADER_PATH /secret/reference")
				}
				return tasking.LearnerVerificationResult{Scope: scope, Passed: false, Summary: "Full verification did not pass.", Diagnostic: "PRIVATE_GRADER_PATH"}, nil
			}
			u := testLocalUI(t, s)
			id := "request-redaction-001"
			w := uiHTTP(u, "POST", "/api/action", uiActionBody(t, s, "complete", "literal-bindings-001", id), nil)
			if w.Code != 202 {
				t.Fatalf("admission: %d", w.Code)
			}
			job := waitUIJob(t, u, id)
			data, _ := json.Marshal(job)
			if bytes.Contains(data, []byte("PRIVATE_GRADER_PATH")) || bytes.Contains(data, []byte("/secret")) {
				t.Fatalf("private error leaked: %s", data)
			}
			if internalError && job.State != "failed" {
				t.Fatal("internal failure presented as success")
			}
			p, err := s.Progress()
			if err != nil || p.Tasks["literal-bindings-001"].Completed {
				t.Fatal("failed completion counted")
			}
		})
	}
}

func TestUIRetainsBoundedJobsAndRestartRejectsOldSession(t *testing.T) {
	s := uiTestService(t, "")
	if _, err := s.Start("literal-bindings-001"); err != nil {
		t.Fatal(err)
	}
	u := testLocalUI(t, s)
	for i := 0; i < 18; i++ {
		id := fmt.Sprintf("request-retention-%03d", i)
		w := uiHTTP(u, "POST", "/api/action", uiActionBody(t, s, "pause", "literal-bindings-001", id), nil)
		if w.Code != 202 {
			t.Fatalf("admission %d: %d", i, w.Code)
		}
		if job := waitUIJob(t, u, id); job.State != "finished" {
			t.Fatalf("pause failed: %+v", job)
		}
	}
	if len(u.jobs) != 16 || len(u.requests) != 16 || len(u.order) != 16 {
		t.Fatal("job retention unbounded")
	}
	if w := uiHTTP(u, "GET", "/api/job?id=request-retention-000", "", nil); w.Code != 404 {
		t.Fatal("old job not evicted")
	}
	before := persistedBytes(t, s)
	restarted := testLocalUI(t, uiTestService(t, s.home))
	if w := uiHTTP(restarted, "GET", "/api/state", "", func(r *http.Request) { r.Header.Set("Authorization", "Bearer "+u.token) }); w.Code != 401 {
		t.Fatal("old session survived restart")
	}
	if w := uiHTTP(restarted, "GET", "/api/job?id=request-retention-017", "", nil); w.Code != 404 {
		t.Fatal("in-memory job survived restart")
	}
	if !bytes.Equal(before, persistedBytes(t, s)) {
		t.Fatal("restart changed progress")
	}
}

func TestUIOpenUsesValidatedWorkspaceAndUnavailableFallback(t *testing.T) {
	s := uiTestService(t, "")
	id := "literal-bindings-001"
	workspace, err := s.Start(id)
	if err != nil {
		t.Fatal(err)
	}
	called := ideTarget{}
	s.launch = func(target ideTarget) error { called = target; return ErrIDEUnavailable }
	u := testLocalUI(t, s)
	request := "request-launcher-0001"
	if w := uiHTTP(u, "POST", "/api/action", uiActionBody(t, s, "open", id, request), nil); w.Code != 202 {
		t.Fatal(w.Code)
	}
	job := waitUIJob(t, u, request)
	if called.Project != filepath.Join(s.home, "workspaces") || called.File != filepath.Join(workspace, "subject", "bindings.go") || job.State != "failed" || !strings.Contains(job.Error, "copy") {
		t.Fatalf("unavailable fallback: %+v target %+v", job, called)
	}
	// Replacing the temporary fixture marker must not invoke the launcher.
	if err := os.Remove(filepath.Join(workspace, learnerWorkspaceMarker)); err != nil {
		t.Fatal(err)
	}
	called = ideTarget{}
	request = "request-launcher-0002"
	if w := uiHTTP(u, "POST", "/api/action", uiActionBody(t, s, "open", id, request), nil); w.Code != 202 {
		t.Fatal(w.Code)
	}
	if job := waitUIJob(t, u, request); job.State != "failed" || called != (ideTarget{}) {
		t.Fatal("invalid workspace launched")
	}
}
