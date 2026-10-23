package subject_test

import (
	"errors"
	"io"
	"strings"
	"sync/atomic"
	subject "task.local/limited-pipe-relay-001/subject"
	"testing"
	"testing/synctest"
)

type destination struct {
	bytes   []byte
	calls   int
	closed  bool
	allowed int
	failure error
}

func (d *destination) Write(p []byte) (int, error) {
	d.calls++
	n := len(p)
	if d.allowed >= 0 {
		n = min(n, max(0, d.allowed-len(d.bytes)))
	}
	d.bytes = append(d.bytes, p[:n]...)
	if d.allowed >= 0 {
		return n, d.failure
	}
	return n, nil
}
func (d *destination) Close() error { d.closed = true; return nil }

type relayResult struct {
	n              int64
	producer, copy error
}
type relayCall struct {
	done      chan relayResult
	endpoints chan io.Writer
	calls     atomic.Int32
}

func startRelay(dst io.Writer, produce func(io.Writer) error, limit int64) *relayCall {
	c := &relayCall{done: make(chan relayResult, 1), endpoints: make(chan io.Writer, 8)}
	go func() {
		n, p, e := subject.Relay(dst, func(w io.Writer) error { c.calls.Add(1); c.endpoints <- w; return produce(w) }, limit)
		c.done <- relayResult{n, p, e}
	}()
	return c
}
func finishRelay(t *testing.T, c *relayCall) relayResult {
	t.Helper()
	synctest.Wait()
	var result relayResult
	select {
	case result = <-c.done:
	default:
		t.Error("Relay did not finish after copy/producer could complete; close both endpoints and join")
		// Capture actual callback endpoint; exposing Close is only a failure-rescue option.
		for len(c.endpoints) > 0 {
			if closer, ok := (<-c.endpoints).(io.Closer); ok {
				_ = closer.Close()
			}
		}
		synctest.Wait()
		select {
		case result = <-c.done:
		default:
			t.Error("Relay incomplete after fixture endpoint rescue")
		}
	}
	if calls := c.calls.Load(); calls != 1 {
		t.Errorf("producer calls=%d; want exactly 1", calls)
	}
	return result
}
func relayValues(t *testing.T, got relayResult, n int64, p, e error, dst *destination, body string) {
	t.Helper()
	if got.n != n {
		t.Errorf("written count=%d; want %d", got.n, n)
	}
	if got.producer != p {
		t.Errorf("producerErr=%v; want exact callback error %v", got.producer, p)
	}
	if got.copy != e {
		t.Errorf("copyErr=%v; want exact copy error %v", got.copy, e)
	}
	if string(dst.bytes) != body {
		t.Errorf("destination bytes=%q; want %q", dst.bytes, body)
	}
	if dst.closed {
		t.Error("caller-owned destination was closed")
	}
}
func completeRelay(t *testing.T, body string, limit int64, terminal error) {
	t.Helper()
	synctest.Test(t, func(t *testing.T) {
		dst := &destination{allowed: -1}
		call := startRelay(dst, func(w io.Writer) error {
			if body != "" {
				n, err := io.WriteString(w, body)
				if n != len(body) || err != nil {
					t.Errorf("complete producer write=(%d,%v); want (%d,nil)", n, err, len(body))
				}
			}
			return terminal
		}, limit)
		wantErr := terminal
		if int64(len(body)) == limit || terminal == io.EOF {
			wantErr = nil
		}
		relayValues(t, finishRelay(t, call), int64(len(body)), terminal, wantErr, dst, body)
	})
}
func TestPublicEOFAndProducerTerminal(t *testing.T) {
	completeRelay(t, "", 9, nil)
	completeRelay(t, "short", 13, nil)
	completeRelay(t, "part", 12, errors.New("producer terminal"))
	completeRelay(t, "exact", 5, errors.New("unprobed terminal at cap"))
}

type writeObservation struct {
	n   int
	err error
}

func interruptedRelay(t *testing.T, limit int64, dst *destination, wantN int64, wantCopy error) {
	t.Helper()
	synctest.Test(t, func(t *testing.T) {
		release := make(chan struct{})
		writes := make(chan writeObservation, 8)
		terminal := errors.New("callback result after interrupted write")
		body := "0123456789abcdefghijklmnopqrstuv"
		call := startRelay(dst, func(w io.Writer) error {
			n, e := io.WriteString(w, body)
			writes <- writeObservation{n, e}
			<-release
			return terminal
		}, limit)
		synctest.Wait()
		var early *relayResult
		select {
		case result := <-call.done:
			early = &result
			t.Error("Relay returned before producer callback release; want joined callback")
		default:
		}
		var observed *writeObservation
		select {
		case w := <-writes:
			observed = &w
		default:
			t.Error("producer write remained blocked after consumer stopped")
		}
		// Release first, then finishRelay may rescue a concrete missing-close endpoint.
		close(release)
		if early != nil {
			call.done <- *early
		}
		got := finishRelay(t, call)
		if observed == nil {
			select {
			case w := <-writes:
				observed = &w
			default:
				t.Error("producer write observation absent after cleanup")
			}
		}
		wantWrite := wantCopy
		if wantWrite == nil {
			wantWrite = io.ErrClosedPipe
		}
		if observed != nil {
			if observed.err != wantWrite {
				t.Errorf("interrupted producer Write error=%v; want exact reader-close error %v", observed.err, wantWrite)
			}
			if int64(observed.n) > limit {
				t.Errorf("pipe accepted %d bytes; limit is %d", observed.n, limit)
			}
			if wantCopy == nil && int64(observed.n) != limit {
				t.Errorf("pipe accepted prefix=%d; want cap %d", observed.n, limit)
			}
		}
		relayValues(t, got, wantN, terminal, wantCopy, dst, body[:int(wantN)])
		if limit == 0 && dst.calls != 0 {
			t.Errorf("destination Write calls=%d at zero limit; want 0", dst.calls)
		}
	})
}
func TestPublicCapAndProducerJoin(t *testing.T) {
	interruptedRelay(t, 4, &destination{allowed: -1}, 4, nil)
	interruptedRelay(t, 0, &destination{allowed: -1}, 0, nil)
}
func TestPublicDestinationFailurePropagation(t *testing.T) {
	marker := errors.New("destination failure")
	interruptedRelay(t, 9, &destination{allowed: 3, failure: marker}, 3, marker)
	interruptedRelay(t, 8, &destination{allowed: 2}, 2, io.ErrShortWrite)
}
func TestPublicBinaryAndIndependentRelay(t *testing.T) {
	completeRelay(t, "a\x00Ω\xff", 16, nil)
	completeRelay(t, strings.Repeat("x", 8), 12, nil)
}
