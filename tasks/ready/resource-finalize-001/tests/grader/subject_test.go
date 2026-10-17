package subject_test

import (
	"errors"
	"io"
	"slices"
	"strings"
	"task.local/resource-finalize-001/subject"
	"testing"
)

func TestStoreFailureCombinations(t *testing.T) {
	for size := 0; size <= 4; size++ {
		for n := 0; n <= size; n++ {
			for _, writeFails := range []bool{false, true} {
				for _, closeFails := range []bool{false, true} {
					data := make([]byte, size)
					for i := range data {
						data[i] = byte(30 + i)
					}
					p := &probeResource{n: n}
					if writeFails {
						p.writeErr = writeSentinel
					}
					if closeFails {
						p.closeErr = closeSentinel
					}
					err := subject.Store(p, data)
					if !slices.Equal(p.events, []string{"write", "close"}) || !slices.Equal(p.data, data) {
						t.Fatalf("size %d n %d: calls/data = %v/%v, want [write close]/%v", size, n, p.events, p.data, data)
					}
					short := !writeFails && n < size
					if errors.Is(err, writeSentinel) != writeFails || errors.Is(err, closeSentinel) != closeFails || errors.Is(err, io.ErrShortWrite) != short {
						t.Errorf("size %d n %d write %t close %t: identities = %t/%t/%t, want %t/%t/%t", size, n, writeFails, closeFails, errors.Is(err, writeSentinel), errors.Is(err, closeSentinel), errors.Is(err, io.ErrShortWrite), writeFails, closeFails, short)
					}
					if writeFails && (err == nil || !strings.Contains(strings.ToLower(err.Error()), "write")) {
						t.Errorf("write context = %v, want operation context", err)
					}
					if !writeFails && !closeFails && !short && err != nil {
						t.Errorf("success = %v, want nil", err)
					}
				}
			}
		}
	}
}
