package subject_test

import (
	"errors"
	"io"
	"strings"
	"testing"
	"testing/synctest"
)

func TestGraderBoundaries(t *testing.T) {
	completeRelay(t, "", 0, errors.New("empty cap terminal"))
	completeRelay(t, "abc", 16, io.EOF)
	completeRelay(t, strings.Repeat("b", 16), 16, errors.New("exact16"))
	interruptedRelay(t, 16, &destination{allowed: -1}, 16, nil)
	marker := errors.New("no output")
	interruptedRelay(t, 1, &destination{allowed: 0, failure: marker}, 0, marker)
}

func TestGraderMultipleWrites(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		dst := &destination{allowed: -1}
		marker := errors.New("multiple write end")
		call := startRelay(dst, func(w io.Writer) error {
			for _, part := range []string{"a", "\x00\xff", "bc"} {
				if _, err := io.WriteString(w, part); err != nil {
					return err
				}
			}
			return marker
		}, 11)
		relayValues(t, finishRelay(t, call), 5, marker, marker, dst, "a\x00\xffbc")
	})
}
