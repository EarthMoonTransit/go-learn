package isolation

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

// RuntimeInfo is the actual queried Go compiler and native target. Collection
// freshness applies to all profiles; only the finite fuzz kernel imposes a
// narrower supported compiler family.
type RuntimeInfo struct {
	Version string `yaml:"version" json:"version"`
	GOOS    string `yaml:"goos" json:"goos"`
	GOARCH  string `yaml:"goarch" json:"goarch"`
}

var platformName = regexp.MustCompile(`^[a-z][a-z0-9]{1,15}$`)

func (info RuntimeInfo) Validate() error {
	if info.Version == "" || len(info.Version) > 160 || strings.TrimSpace(info.Version) != info.Version || !platformName.MatchString(info.GOOS) || !platformName.MatchString(info.GOARCH) {
		return fmt.Errorf("incomplete actual Go runtime identity")
	}
	return nil
}

func ParseRuntimeInfo(data []byte) (RuntimeInfo, error) {
	var values struct{ GOVERSION, GOOS, GOARCH string }
	if len(data) > 64<<10 {
		return RuntimeInfo{}, fmt.Errorf("Go runtime query exceeds capture bound")
	}
	if err := json.Unmarshal(data, &values); err != nil {
		return RuntimeInfo{}, fmt.Errorf("decode actual Go runtime: %w", err)
	}
	info := RuntimeInfo{Version: values.GOVERSION, GOOS: values.GOOS, GOARCH: values.GOARCH}
	if err := info.Validate(); err != nil {
		return RuntimeInfo{}, err
	}
	return info, nil
}
