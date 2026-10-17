package subject_test

import (
	"errors"
	"fmt"
	"io"
	"slices"
	"strings"
	"task.local/resource-finalize-001/subject"
	"testing"
)

var writeSentinel = errors.New("origin-a")
var closeSentinel = errors.New("origin-b")

type probeResource struct {
	n                  int
	writeErr, closeErr error
	events             []string
	data               []byte
	panicValue         any
}

func (p *probeResource) Write(data []byte) (int, error) {
	p.events = append(p.events, "write")
	p.data = append([]byte(nil), data...)
	if p.panicValue != nil {
		panic(p.panicValue)
	}
	n := p.n
	if n > len(data) {
		n = len(data)
	}
	return n, p.writeErr
}
func (p *probeResource) Close() error { p.events = append(p.events, "close"); return p.closeErr }

func TestStorePublicContract(t *testing.T) {
	for _, c := range []struct {
		name               string
		data               []byte
		n                  int
		writeErr, closeErr error
	}{
		{"success", []byte("abc"), 3, nil, nil},
		{"write-only", []byte("abc"), 0, writeSentinel, nil},
		{"close-only", []byte("abc"), 3, nil, closeSentinel},
		{"both", []byte("abc"), 1, writeSentinel, closeSentinel},
		{"short", []byte("abc"), 1, nil, nil},
		{"short-and-close", []byte("abc"), 1, nil, closeSentinel},
		{"empty", nil, 0, nil, nil},
		{"empty-and-errors", []byte{}, 0, writeSentinel, closeSentinel},
	} {
		t.Run(c.name, func(t *testing.T) {
			p := &probeResource{n: c.n, writeErr: c.writeErr, closeErr: c.closeErr}
			err := subject.Store(p, c.data)
			if !slices.Equal(p.events, []string{"write", "close"}) {
				t.Errorf("calls/order = %v, want [write close]", p.events)
			}
			if !slices.Equal(p.data, c.data) {
				t.Errorf("Write data = %v, want %v", p.data, c.data)
			}
			wantWrite, wantClose := c.writeErr != nil, c.closeErr != nil
			wantShort := c.writeErr == nil && c.n < len(c.data)
			if got := errors.Is(err, writeSentinel); got != wantWrite {
				t.Errorf("write identity = %t, want %t; error %v", got, wantWrite, err)
			}
			if got := errors.Is(err, closeSentinel); got != wantClose {
				t.Errorf("close identity = %t, want %t; error %v", got, wantClose, err)
			}
			if got := errors.Is(err, io.ErrShortWrite); got != wantShort {
				t.Errorf("short-write identity = %t, want %t; error %v", got, wantShort, err)
			}
			if !wantWrite && !wantClose && !wantShort && err != nil {
				t.Errorf("success error = %v, want nil", err)
			}
			if wantWrite && (err == nil || !strings.Contains(strings.ToLower(err.Error()), "write")) {
				t.Errorf("write context = %v, want error text naming write", err)
			}
		})
	}
}

func storePanic(call func()) (value any) { defer func() { value = recover() }(); call(); return nil }
func TestStorePreservesWritePanicAndCloses(t *testing.T) {
	token := &struct{ code int }{19}
	p := &probeResource{panicValue: token}
	got := storePanic(func() { subject.Store(p, []byte("abc")) })
	if got != token {
		t.Errorf("write panic = %v, want original %v", got, token)
	}
	if !slices.Equal(p.events, []string{"write", "close"}) {
		t.Errorf("panic calls/order = %v, want [write close]", p.events)
	}
}

func TestStoreKeepsDistinctCausesWithEqualText(t *testing.T) {
	writeCause := errors.New("same")
	closeCause := errors.New("same")
	writeWrapper := fmt.Errorf("operation: %w", writeCause)
	closeWrapper := fmt.Errorf("operation: %w", closeCause)
	p := &probeResource{n: 0, writeErr: writeWrapper, closeErr: closeWrapper}
	err := subject.Store(p, []byte("abc"))
	for _, check := range []struct {
		name   string
		target error
	}{
		{"write wrapper", writeWrapper}, {"close wrapper", closeWrapper},
		{"write cause", writeCause}, {"close cause", closeCause},
	} {
		if got := errors.Is(err, check.target); !got {
			t.Errorf("%s identity = %t, want true; error %v", check.name, got, err)
		}
	}
	if !slices.Equal(p.events, []string{"write", "close"}) {
		t.Errorf("equal-text calls/order = %v, want [write close]", p.events)
	}
}
