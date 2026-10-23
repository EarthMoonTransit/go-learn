package subject_test

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strings"
	subject "task.local/buffered-word-stream-001/subject"
	"testing"
)

type wordDestination struct {
	data          []byte
	calls, closes int
	quota         int
	failure       error
	failed        bool
}

func (d *wordDestination) Write(p []byte) (int, error) {
	d.calls++
	n := len(p)
	if d.quota >= 0 && n > d.quota {
		n = d.quota
	}
	d.data = append(d.data, p[:n]...)
	if d.quota >= 0 {
		d.quota -= n
		d.failed = true
		return n, d.failure
	}
	return n, nil
}
func (d *wordDestination) Close() error { d.closes++; return nil }

type wordSource struct {
	data                            []byte
	terminal                        error
	together                        bool
	step                            int
	reads, closes, postFailureReads int
	destination                     *wordDestination
	atTerminal                      func()
	terminalSeen                    bool
}

func (s *wordSource) Read(p []byte) (int, error) {
	s.reads++
	if s.destination != nil && s.destination.failed {
		s.postFailureReads++
	}
	terminal := s.terminal
	if terminal == nil {
		terminal = io.EOF
	}
	if len(s.data) == 0 {
		if !s.terminalSeen {
			s.terminalSeen = true
			if s.atTerminal != nil {
				s.atTerminal()
			}
		}
		return 0, terminal
	}
	n := len(s.data)
	if n > len(p) {
		n = len(p)
	}
	if s.step > 0 && n > s.step {
		n = s.step
	}
	copy(p, s.data[:n])
	s.data = s.data[n:]
	if s.together && len(s.data) == 0 {
		return n, terminal
	}
	return n, nil
}
func (s *wordSource) Close() error { s.closes++; return nil }
func openWordStreams(t *testing.T, s *wordSource, d *wordDestination) {
	t.Helper()
	if s.closes != 0 || d.closes != 0 {
		t.Errorf("stream ownership: source closes=%d destination closes=%d; want both 0", s.closes, d.closes)
	}
}
func normalWords(t *testing.T, input, want string, step int) {
	t.Helper()
	dst := &wordDestination{quota: -1}
	src := &wordSource{data: []byte(input), step: step}
	err := subject.WriteWords(dst, src)
	if err != nil {
		t.Errorf("word normalization returned %v; want nil", err)
	}
	if string(dst.data) != want {
		t.Errorf("normalized words=%q; want %q", dst.data, want)
	}
	if want == "" && dst.calls != 0 {
		t.Errorf("empty word stream writes=%d; want 0", dst.calls)
	}
	openWordStreams(t, src, dst)
}
func TestPublicWordOrderAndFinalToken(t *testing.T) {
	normalWords(t, "  alpha\tbeta\nrepeat repeat\u2003λ-last", "alpha\nbeta\nrepeat\nrepeat\nλ-last\n", 3)
	normalWords(t, "", "", 1)
	normalWords(t, " \n\t\u2003", "", 2)
}
func TestPublicTokenBounds(t *testing.T) {
	normalWords(t, strings.Repeat("q", 64), strings.Repeat("q", 64)+"\n", 5)
	normalWords(t, strings.Repeat("λ", 32), strings.Repeat("λ", 32)+"\n", 7)
	dst := &wordDestination{quota: -1}
	src := &wordSource{data: []byte("ok " + strings.Repeat("x", 1024)), step: 9}
	err := subject.WriteWords(dst, src)
	if err == nil {
		t.Error("1024-byte unbroken word must produce a scanning error")
	}
	if string(dst.data) != "ok\n" {
		t.Errorf("oversize accepted prefix=%q; want %q", dst.data, "ok\n")
	}
	openWordStreams(t, src, dst)
}
func TestPublicReaderErrorKeepsFinalWord(t *testing.T) {
	marker := errors.New("source terminal marker")
	dst := &wordDestination{quota: -1}
	src := &wordSource{data: []byte("ok tail"), terminal: marker, together: true}
	err := subject.WriteWords(dst, src)
	if err != marker {
		t.Errorf("reader terminal error=%v; want original marker", err)
	}
	if string(dst.data) != "ok\ntail\n" {
		t.Errorf("terminal accepted words=%q; want %q", dst.data, "ok\ntail\n")
	}
	openWordStreams(t, src, dst)
}
func TestPublicOwnedBufferBeforeFinalFlush(t *testing.T) {
	dst := &wordDestination{quota: -1}
	src := &wordSource{data: []byte("aa b ")}
	src.atTerminal = func() {
		if dst.calls != 0 {
			t.Errorf("before final flush: destination writes=%d; want 0 for output below16 bytes", dst.calls)
		}
	}
	if err := subject.WriteWords(dst, src); err != nil {
		t.Errorf("short stream error=%v; want nil", err)
	}
	if string(dst.data) != "aa\nb\n" {
		t.Errorf("flushed short words=%q; want %q", dst.data, "aa\nb\n")
	}
	openWordStreams(t, src, dst)
}
func TestPublicFreshBufferAroundBufferedDestination(t *testing.T) {
	raw := &wordDestination{quota: -1}
	caller := bufio.NewWriterSize(raw, 8192)
	if _, err := caller.WriteString("pending|"); err != nil {
		t.Fatal(err)
	}
	before := caller.Buffered()
	src := &wordSource{data: []byte("a bb")}
	err := subject.WriteWords(caller, src)
	if err != nil {
		t.Errorf("nested buffer error=%v; want nil", err)
	}
	if len(raw.data) != 0 {
		t.Errorf("caller buffer was flushed: raw bytes=%q; want empty", raw.data)
	}
	if got := caller.Buffered(); got != before+5 {
		t.Errorf("caller pending bytes=%d; want %d after owned flush", got, before+5)
	}
	if err := caller.Flush(); err != nil {
		t.Fatal(err)
	}
	if string(raw.data) != "pending|a\nbb\n" {
		t.Errorf("caller final bytes=%q; want %q", raw.data, "pending|a\nbb\n")
	}
	if src.closes != 0 {
		t.Errorf("nested source closes=%d; want 0", src.closes)
	}
}
func flushFailure(t *testing.T, readerFails, shortNil bool) {
	t.Helper()
	writeMarker := errors.New("destination marker")
	readMarker := errors.New("competing source marker")
	dst := &wordDestination{quota: 2, failure: writeMarker}
	src := &wordSource{data: []byte("abc tail"), together: true}
	want := error(writeMarker)
	if readerFails {
		src.terminal = readMarker
	}
	if shortNil {
		dst.failure = nil
		want = io.ErrShortWrite
	}
	err := subject.WriteWords(dst, src)
	if err != want {
		t.Errorf("final flush priority: error=%v; want exact %v", err, want)
	}
	if !bytes.Equal(dst.data, []byte("ab")) {
		t.Errorf("failed flush accepted bytes=%q; want %q", dst.data, "ab")
	}
	openWordStreams(t, src, dst)
}
func TestPublicFinalFlushErrorPriority(t *testing.T) {
	flushFailure(t, false, false)
	flushFailure(t, true, false)
	flushFailure(t, true, true)
}
func TestPublicNoReadsAfterObservedWriteFailure(t *testing.T) {
	marker := errors.New("buffer drain failed")
	dst := &wordDestination{quota: 3, failure: marker}
	src := &wordSource{data: []byte(strings.Repeat("word ", 2000)), step: 7, destination: dst}
	err := subject.WriteWords(dst, src)
	if err != marker {
		t.Errorf("buffered write failure=%v; want original marker", err)
	}
	if !dst.failed {
		t.Error("finite output never attempted destination write")
	}
	if src.postFailureReads != 0 {
		t.Errorf("source reads after observed buffered write failure=%d; want 0", src.postFailureReads)
	}
	if string(dst.data) != "wor" {
		t.Errorf("failed buffered write prefix=%q; want %q", dst.data, "wor")
	}
	openWordStreams(t, src, dst)
}
