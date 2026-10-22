package subject_test

import (
	"errors"
	"testing"
)

func TestBinaryChunksAndFreshErrors(t *testing.T) {
	checkTransfer(t, "\x00Ω\xff\r\n", nil)
	checkTransfer(t, "", errors.New("empty producer failure"))
	checkTransfer(t, "abcdefghijklmnop", errors.New("end after bytes"))
}
