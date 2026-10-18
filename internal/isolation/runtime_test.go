package isolation

import (
	"strings"
	"testing"
)

func TestRuntimeMetadataDoesNotNarrowExistingProfileSDKs(t *testing.T) {
	for _, version := range []string{"go1.27.0", "go1.28.1", "devel go1.28-test"} {
		info, err := ParseRuntimeInfo([]byte(`{"GOVERSION":"` + version + `","GOOS":"linux","GOARCH":"amd64"}`))
		if err != nil || info.Version != version {
			t.Fatalf("metadata rejected compiler identity %s: %+v %v", version, info, err)
		}
	}
	for _, raw := range []string{`{}`, `{"GOVERSION":"go1.27.0","GOOS":"linux"}`, `{"GOVERSION":"go1.27.0","GOOS":"../linux","GOARCH":"amd64"}`, strings.Repeat("x", 64<<10+1)} {
		if _, err := ParseRuntimeInfo([]byte(raw)); err == nil {
			t.Fatalf("accepted incomplete runtime query %.150s", raw)
		}
	}
}
