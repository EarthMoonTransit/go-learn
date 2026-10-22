package subject_test

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"testing"

	subject "task.local/bounded-record-read-001/subject"
)

// Every source is finite, keeps unread bytes, and obeys len(p).
// Its observations diagnose contract violations without panicking.
type publicRecordSource struct {
	data                                        []byte
	at, calls, closes, afterTerminal, oversized int
	maxChunk, emptyReads, splitAt               int
	terminal                                    error
	attachTerminal, terminalReported            bool
}

func (s *publicRecordSource) Read(p []byte) (int, error) {
	s.calls++
	if len(p) > 3-s.at%3 {
		s.oversized++
	}
	if s.terminalReported {
		s.afterTerminal++
		return 0, s.terminal
	}
	if s.emptyReads > 0 {
		s.emptyReads--
		return 0, nil
	}
	if s.at == len(s.data) {
		s.terminalReported = true
		return 0, s.terminal
	}
	n := len(p)
	if n > len(s.data)-s.at {
		n = len(s.data) - s.at
	}
	if s.maxChunk > 0 && n > s.maxChunk {
		n = s.maxChunk
	}
	if s.splitAt > s.at && s.at+n > s.splitAt {
		n = s.splitAt - s.at
	}
	copy(p[:n], s.data[s.at:s.at+n])
	s.at += n
	if s.attachTerminal && s.at == len(s.data) {
		s.terminalReported = true
		return n, s.terminal
	}
	return n, nil
}
func (s *publicRecordSource) Close() error { s.closes++; return nil }

func publicSource(data []byte, cause error, attach bool) *publicRecordSource {
	if cause == nil {
		cause = io.EOF
	}
	return &publicRecordSource{data: append([]byte(nil), data...), terminal: cause, attachTerminal: attach}
}

func publicNilValue(value any) bool {
	if value == nil {
		return true
	}
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return v.IsNil()
	}
	return false
}

func checkPublicResult(t *testing.T, got []subject.Record, err error, want []subject.Record, cause error) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("accepted records = %v; want %v (length %d)", got, want, len(want))
	} else {
		for i := range want {
			if got[i] != want[i] {
				t.Errorf("record[%d] = %#v; want %#v", i, got[i], want[i])
			}
		}
	}
	if cause == nil {
		if err != nil {
			t.Errorf("success error has dynamic type %T; want actual nil", err)
		}
		return
	}
	if err == nil {
		t.Errorf("failure error = nil; want RecordFailure with cause %q", cause.Error())
		return
	}
	if publicNilValue(err) {
		t.Errorf("failure error has nil dynamic value of type %T; want nonnil RecordFailure", err)
		return
	}
	failure, ok := err.(subject.RecordFailure)
	if !ok {
		t.Errorf("failure type = %T; want RecordFailure for recovery", err)
		return
	}
	if publicNilValue(failure) {
		t.Errorf("RecordFailure has nil dynamic value; want usable recovery state")
		return
	}
	if gotIndex := failure.RecordIndex(); gotIndex != len(want) {
		t.Errorf("RecordIndex() = %d; want accepted-prefix boundary %d", gotIndex, len(want))
	}
	if gotCause := failure.Unwrap(); gotCause != cause {
		t.Errorf("Unwrap() = %v (%T); want exact cause %v (%T)", gotCause, gotCause, cause, cause)
	}
	if !errors.Is(err, cause) {
		t.Errorf("errors.Is(failure, selected cause) = false; want true")
	}
	var extracted subject.RecordFailure
	if !errors.As(err, &extracted) {
		t.Errorf("errors.As to RecordFailure = false; want true")
	}
	wantText := fmt.Sprintf("record %d: %s", len(want), cause.Error())
	if gotText := err.Error(); gotText != wantText {
		t.Errorf("Error() = %q; want %q", gotText, wantText)
	}
}

func checkPublicSource(t *testing.T, source *publicRecordSource, bytes int) {
	t.Helper()
	if source.at != bytes {
		t.Errorf("consumed bytes = %d; want %d, leaving later bytes unread", source.at, bytes)
	}
	if source.oversized != 0 {
		t.Errorf("requests exceeding missing current-record bytes = %d; want 0", source.oversized)
	}
	if source.afterTerminal != 0 {
		t.Errorf("reads after terminal error = %d; want 0", source.afterTerminal)
	}
	if source.closes != 0 {
		t.Errorf("input Close calls = %d; want 0 (caller owns input)", source.closes)
	}
}

func TestPublicZeroCount(t *testing.T) {
	source := publicSource([]byte{1, 3, 7}, nil, false)
	got, err := subject.ReadRecords(source, 0)
	checkPublicResult(t, got, err, nil, nil)
	checkPublicSource(t, source, 0)
	if source.calls != 0 {
		t.Errorf("zero-count Read calls = %d; want 0", source.calls)
	}
}

func TestPublicBoundedPrefix(t *testing.T) {
	for _, chunk := range []int{0, 1, 2} {
		t.Run(fmt.Sprintf("chunk-%d", chunk), func(t *testing.T) {
			source := publicSource([]byte{1, 1, 2, 1, 254, 0, 99, 98, 97}, nil, false)
			source.maxChunk = chunk
			source.emptyReads = 1
			got, err := subject.ReadRecords(source, 2)
			checkPublicResult(t, got, err, []subject.Record{{Value: 0x0102}, {Value: 0xfe00}}, nil)
			checkPublicSource(t, source, 6)
		})
	}
}

func TestPublicTerminalPolicy(t *testing.T) {
	transport := errors.New("reader stopped")
	wrappedEOF := fmt.Errorf("reader wrapped end: %w", io.EOF)
	one := []subject.Record{{Value: 9}}
	initialVersion := subject.ErrVersion
	cases := []struct {
		name     string
		data     []byte
		count    int
		terminal error
		attach   bool
		want     []subject.Record
		cause    error
		consumed int
	}{
		{"empty-eof", nil, 1, io.EOF, false, nil, io.ErrUnexpectedEOF, 0},
		{"one-byte-eof", []byte{1}, 1, io.EOF, true, nil, io.ErrUnexpectedEOF, 1},
		{"incomplete-invalid-version", []byte{7, 2}, 1, io.EOF, true, nil, io.ErrUnexpectedEOF, 2},
		{"boundary-eof", []byte{1, 0, 9}, 2, io.EOF, false, one, io.ErrUnexpectedEOF, 3},
		{"valid-intermediate-eof", []byte{1, 0, 9}, 2, io.EOF, true, one, io.ErrUnexpectedEOF, 3},
		{"partial-next-eof", []byte{1, 0, 9, 1, 2}, 2, io.EOF, true, one, io.ErrUnexpectedEOF, 5},
		{"complete-final-eof", []byte{1, 0, 9, 1, 0, 4}, 2, io.EOF, true, []subject.Record{{Value: 9}, {Value: 4}}, nil, 6},
		{"complete-final-error", []byte{1, 0, 9}, 1, transport, true, one, transport, 3},
		{"complete-intermediate-error", []byte{1, 0, 9}, 2, transport, true, one, transport, 3},
		{"partial-next-error", []byte{1, 0, 9, 1, 3}, 2, transport, true, one, transport, 5},
		{"empty-error", nil, 1, transport, false, nil, transport, 0},
		{"bad-version-with-tail", []byte{1, 0, 9, 7, 0, 4, 1, 0, 3}, 3, io.EOF, false, one, initialVersion, 6},
		{"bad-version-eof", []byte{1, 0, 9, 7, 0, 4}, 2, io.EOF, true, one, initialVersion, 6},
		{"bad-version-reader-error", []byte{1, 0, 9, 7, 0, 4}, 2, transport, true, one, transport, 6},
		{"wrapped-eof-is-reader-error", []byte{1, 0, 9}, 1, wrappedEOF, true, one, wrappedEOF, 3},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			source := publicSource(tc.data, tc.terminal, tc.attach)
			got, err := subject.ReadRecords(source, tc.count)
			checkPublicResult(t, got, err, tc.want, tc.cause)
			checkPublicSource(t, source, tc.consumed)
			if subject.ErrVersion != initialVersion {
				t.Errorf("ErrVersion identity changed; want original supplied sentinel")
			}
		})
	}
}

func TestPublicRetainedFailures(t *testing.T) {
	firstCause := errors.New("first stopped")
	firstSource := publicSource(nil, firstCause, false)
	first, firstError := subject.ReadRecords(firstSource, 1)
	checkPublicResult(t, first, firstError, nil, firstCause)
	checkPublicSource(t, firstSource, 0)

	secondCause := errors.New("second stopped")
	secondSource := publicSource([]byte{1, 0, 2, 1, 0, 6, 1}, secondCause, true)
	second, secondError := subject.ReadRecords(secondSource, 3)
	checkPublicResult(t, second, secondError, []subject.Record{{Value: 2}, {Value: 6}}, secondCause)
	checkPublicSource(t, secondSource, 7)
	// Re-observe the retained first failure after another call created a failure.
	checkPublicResult(t, first, firstError, nil, firstCause)
}
