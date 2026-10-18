// Package fuzznative admits and observes the finite native hex seed profile.
// It uses disposable workspaces under the repository's existing isolation model.
package fuzznative

import (
	"time"

	"go-practice-bank/internal/fuzznative/frame"
	"go-practice-bank/internal/fuzznative/protocol"
)

type Identity = frame.Identity
type Runtime = protocol.Runtime

type Request struct {
	Identity       Identity
	GoVersion      string
	Implementation map[string]string
	Sources        map[string]string
	Timeout        time.Duration
	Race           bool
	Vet            bool
}

type Result struct {
	Valid           bool
	Detected        bool
	Canonical       string
	Diagnostic      string
	Phase           string
	Vetted          bool
	RaceUnsupported bool
	Callbacks       int
	Calls           int
	Assertions      int
}

type Metrics struct {
	Workspaces int
	GoRuns     int
	HelperRuns int
	NativeRuns int
}
