package learner

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"embed"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"net"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

//go:embed ui/index.html ui/app.js ui/app.css
var uiAssets embed.FS

var requestIDPattern = regexp.MustCompile(`^[A-Za-z0-9_-]{16,80}$`)

type uiRequest struct {
	Action    string `json:"action"`
	TaskID    string `json:"task_id"`
	Revision  string `json:"revision"`
	RequestID string `json:"request_id"`
}
type uiJob struct {
	ID     string        `json:"id"`
	Action string        `json:"action"`
	TaskID string        `json:"task_id"`
	State  string        `json:"state"`
	Result *ActionResult `json:"result,omitempty"`
	Error  string        `json:"error,omitempty"`
}
type localUI struct {
	service  *Service
	host     string
	token    string
	page     *template.Template
	mu       sync.Mutex
	stopping bool
	active   string
	jobs     map[string]uiJob
	requests map[string]uiRequest
	order    []string
	workers  sync.WaitGroup
}

func newLocalUI(service *Service, host string) (*localUI, error) {
	secret := make([]byte, 32)
	if _, err := rand.Read(secret); err != nil {
		return nil, err
	}
	page, err := template.ParseFS(uiAssets, "ui/index.html")
	if err != nil {
		return nil, err
	}
	return &localUI{service: service, host: host, token: hex.EncodeToString(secret), page: page, jobs: map[string]uiJob{}, requests: map[string]uiRequest{}}, nil
}

// ServeUI is foreground and drains any admitted verification before returning.
// Jobs are process-local; a new run has a new token and no previous job records.
func ServeUI(ctx context.Context, service *Service, port int, output io.Writer) error {
	if port < 0 || port > 65535 {
		return errors.New("UI port must be between 0 and 65535")
	}
	listener, err := net.Listen("tcp4", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		return err
	}
	ui, err := newLocalUI(service, listener.Addr().String())
	if err != nil {
		listener.Close()
		return err
	}
	server := &http.Server{Handler: ui, ReadHeaderTimeout: 5 * time.Second, ReadTimeout: 10 * time.Second, WriteTimeout: 30 * time.Second, IdleTimeout: 30 * time.Second, MaxHeaderBytes: 16 << 10}
	fmt.Fprintf(output, "Go Learn UI: http://%s/#%s\nRoot: %s\nHome: %s\nKeep this process running. Press Ctrl+C to stop; an active check will finish first.\n", ui.host, ui.token, service.root, service.home)
	served := make(chan error, 1)
	go func() { served <- server.Serve(listener) }()
	select {
	case <-ctx.Done():
	case err = <-served:
		if errors.Is(err, http.ErrServerClosed) {
			err = nil
		}
	}
	ui.beginShutdown()
	fmt.Fprintln(output, "Stopping UI; waiting for any active operation to finish.")
	// Stop listeners and HTTP traffic promptly, but don't orphan the independent
	// worker: it owns the operation lease until verification and persistence finish.
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	shutdownErr := server.Shutdown(shutdownCtx)
	cancel()
	if shutdownErr != nil {
		_ = server.Close()
	}
	ui.workers.Wait()
	return err
}

func (u *localUI) beginShutdown() {
	u.mu.Lock()
	u.stopping = true
	u.mu.Unlock()
}
func writeUIJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
func uiError(w http.ResponseWriter, status int, message string) {
	writeUIJSON(w, status, map[string]string{"error": message})
}
func safeUIError(err error) string {
	switch {
	case errors.Is(err, ErrOperationBusy):
		return "Another learner command is active, or a lock needs manual recovery. See docs/LEARNER_CLI.md."
	case errors.Is(err, ErrStaleView):
		return "Learner state changed. No action was run. Refresh and try again."
	case errors.Is(err, ErrIDEUnavailable):
		return ErrIDEUnavailable.Error()
	default:
		return "The operation could not finish. Check your workspace and use the matching learn CLI for details."
	}
}
func (u *localUI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Referrer-Policy", "no-referrer")
	w.Header().Set("Content-Security-Policy", "default-src 'none'; script-src 'self'; style-src 'self'; connect-src 'self'; base-uri 'none'; form-action 'none'; frame-ancestors 'none'")
	if r.Host != u.host {
		uiError(w, http.StatusForbidden, "Unrecognized local host.")
		return
	}
	if len(r.RequestURI) > 2048 {
		uiError(w, http.StatusRequestURITooLong, "Request is too long.")
		return
	}
	origin := "http://" + u.host
	if (r.Header.Get("Origin") != "" && r.Header.Get("Origin") != origin) || r.Header.Get("Sec-Fetch-Site") == "cross-site" {
		uiError(w, http.StatusForbidden, "Cross-origin requests are not allowed.")
		return
	}
	if r.Method == http.MethodGet && r.URL.RawQuery == "" {
		switch r.URL.Path {
		case "/":
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			_ = u.page.Execute(w, nil)
			return
		case "/app.js", "/app.css":
			content, err := uiAssets.ReadFile("ui" + r.URL.Path)
			if err != nil {
				uiError(w, 500, "Asset unavailable.")
				return
			}
			if r.URL.Path == "/app.js" {
				w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
			} else {
				w.Header().Set("Content-Type", "text/css; charset=utf-8")
			}
			_, _ = w.Write(content)
			return
		}
	}
	if !strings.HasPrefix(r.URL.Path, "/api/") {
		uiError(w, 404, "Not found.")
		return
	}
	authorization := r.Header.Get("Authorization")
	provided := strings.TrimPrefix(authorization, "Bearer ")
	if !strings.HasPrefix(authorization, "Bearer ") || subtle.ConstantTimeCompare([]byte(provided), []byte(u.token)) != 1 {
		uiError(w, 401, "Reopen the full URL printed by the current learn ui process.")
		return
	}
	if r.Method == http.MethodGet {
		switch r.URL.Path {
		case "/api/state":
			if r.URL.RawQuery != "" {
				uiError(w, 400, "Unexpected query.")
				return
			}
			course, err := u.service.CourseView()
			if err != nil {
				uiError(w, 500, safeUIError(err))
				return
			}
			u.mu.Lock()
			var job *uiJob
			if u.active != "" {
				copy := u.jobs[u.active]
				job = &copy
			}
			stopping := u.stopping
			u.mu.Unlock()
			writeUIJSON(w, 200, struct {
				Course   CourseView `json:"course"`
				Job      *uiJob     `json:"job"`
				Stopping bool       `json:"stopping"`
			}{course, job, stopping})
			return
		case "/api/topics":
			if r.URL.RawQuery != "" {
				uiError(w, 400, "Unexpected query.")
				return
			}
			topics, err := u.service.Topics()
			if err != nil {
				uiError(w, 500, safeUIError(err))
				return
			}
			writeUIJSON(w, 200, topics)
			return
		case "/api/theory":
			query := r.URL.Query()
			id := query.Get("id")
			if len(query) != 1 || len(query["id"]) != 1 || !topicIDPattern.MatchString(id) {
				uiError(w, 400, "Invalid topic ID.")
				return
			}
			view, err := u.service.TheoryView(id)
			if err != nil {
				uiError(w, 404, "Theory unavailable. Use learn theory for details.")
				return
			}
			writeUIJSON(w, 200, view)
			return
		case "/api/task":
			query := r.URL.Query()
			id := query.Get("id")
			if len(query) != 1 || len(query["id"]) != 1 || !learnerTaskIDPattern.MatchString(id) {
				uiError(w, 400, "Invalid task ID.")
				return
			}
			view, err := u.service.TaskView(id)
			if err != nil {
				uiError(w, 400, safeUIError(err))
				return
			}
			writeUIJSON(w, 200, view)
			return
		case "/api/job":
			query := r.URL.Query()
			id := query.Get("id")
			if len(query) != 1 || len(query["id"]) != 1 || !requestIDPattern.MatchString(id) {
				uiError(w, 400, "Invalid operation ID.")
				return
			}
			u.mu.Lock()
			job, ok := u.jobs[id]
			u.mu.Unlock()
			if !ok {
				uiError(w, 404, "Operation is no longer available. Refresh progress; jobs do not survive a UI restart.")
				return
			}
			writeUIJSON(w, 200, job)
			return
		}
		uiError(w, 404, "Not found.")
		return
	}
	if r.URL.Path != "/api/action" {
		uiError(w, 404, "Not found.")
		return
	}
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		uiError(w, 405, "Use POST for actions.")
		return
	}
	if r.Header.Get("Origin") != origin {
		uiError(w, 403, "Actions require the local UI origin.")
		return
	}
	if r.URL.RawQuery != "" || r.Header.Get("Content-Type") != "application/json" {
		uiError(w, 400, "Use a JSON action request.")
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 4096)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var request uiRequest
	if err := decoder.Decode(&request); err != nil {
		uiError(w, 400, "Invalid action request.")
		return
	}
	var trailing any
	if err := decoder.Decode(&trailing); err != io.EOF {
		uiError(w, 400, "Expected one action request.")
		return
	}
	if !requestIDPattern.MatchString(request.RequestID) || !learnerTaskIDPattern.MatchString(request.TaskID) || len(request.Revision) != 64 {
		uiError(w, 400, "Invalid action identity. Refresh and try again.")
		return
	}
	switch request.Action {
	case "start", "resume", "check", "complete", "hint", "pause", "open":
	default:
		uiError(w, 400, "Unknown action.")
		return
	}
	u.admit(w, request)
}

func (u *localUI) admit(w http.ResponseWriter, request uiRequest) {
	u.mu.Lock()
	if u.stopping {
		u.mu.Unlock()
		uiError(w, 503, "UI is stopping; no new action was started.")
		return
	}
	if previous, ok := u.requests[request.RequestID]; ok {
		job := u.jobs[request.RequestID]
		u.mu.Unlock()
		if previous != request {
			uiError(w, 409, "Operation ID was reused with different input.")
			return
		}
		writeUIJSON(w, 200, job)
		return
	}
	if u.active != "" {
		u.mu.Unlock()
		uiError(w, 409, "An operation is already running. No second action was started.")
		return
	}
	if len(u.order) >= 16 {
		oldest := u.order[0]
		delete(u.jobs, oldest)
		delete(u.requests, oldest)
		u.order = u.order[1:]
	}
	job := uiJob{ID: request.RequestID, Action: request.Action, TaskID: request.TaskID, State: "running"}
	u.requests[request.RequestID] = request
	u.jobs[request.RequestID] = job
	u.order = append(u.order, request.RequestID)
	u.active = request.RequestID
	u.workers.Add(1)
	u.mu.Unlock()
	go u.run(request)
	writeUIJSON(w, 202, job)
}
func (u *localUI) run(request uiRequest) {
	defer u.workers.Done()
	result, err := u.service.PerformAction(request.Action, request.TaskID, request.Revision)
	// Complete never returns public compiler/test diagnostics through HTTP.
	if result.Verification != nil && request.Action == "complete" {
		result.Verification.Diagnostic = ""
	}
	u.mu.Lock()
	job := u.jobs[request.RequestID]
	if err != nil {
		job.State = "failed"
		job.Error = safeUIError(err)
	} else {
		job.State = "finished"
		job.Result = &result
	}
	u.jobs[request.RequestID] = job
	u.active = ""
	u.mu.Unlock()
}
