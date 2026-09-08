package subject

import (
	"strings"
)

func Classify(identifier string, version int) (bool, bool) {
	identifier = strings.TrimSpace(identifier)
	listed := identifier == "alpha" || identifier == "beta"
	return listed, version >= 18
}
