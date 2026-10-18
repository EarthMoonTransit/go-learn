package protocol

import (
	"fmt"
	"regexp"
	"runtime"
)

// Runtime names the compiler's actual version and target, rather than those of
// an installed caller which may have been built by another compiler.
type Runtime struct {
	Version string `json:"version"`
	GOOS    string `json:"goos"`
	GOARCH  string `json:"goarch"`
}

var platformPattern = regexp.MustCompile(`^[a-z][a-z0-9]{1,15}$`)

func (r Runtime) Validate() error {
	if !runtimePattern.MatchString(r.Version) || !platformPattern.MatchString(r.GOOS) || !platformPattern.MatchString(r.GOARCH) {
		return fmt.Errorf("unsupported native fuzz runtime %q/%s/%s", r.Version, r.GOOS, r.GOARCH)
	}
	return nil
}

func CurrentRuntime() Runtime {
	return Runtime{Version: runtime.Version(), GOOS: runtime.GOOS, GOARCH: runtime.GOARCH}
}

type Export struct {
	Path   string `json:"path"`
	SHA256 string `json:"sha256"`
}
