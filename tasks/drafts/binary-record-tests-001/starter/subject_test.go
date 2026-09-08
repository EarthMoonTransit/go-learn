package subject_test

import (
	"slices"
	"task.local/binary-record-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	r := subject.Record{Code: 0x1234, Count: 0x89abcdef}
	data, e := r.MarshalBinary()
	if e != nil || !slices.Equal(data, []byte{0x12, 0x34, 0x89, 0xab, 0xcd, 0xef}) {
		t.Errorf("wire got%x,%v", data, e)
	}
	var got subject.Record
	if e = got.UnmarshalBinary([]byte{0x12, 0x34, 0x89, 0xab, 0xcd, 0xef}); e != nil || got != r {
		t.Errorf("decode got%+v,%v", got, e)
	}
}

// Add cases for the remaining requirements in TASK.md.
