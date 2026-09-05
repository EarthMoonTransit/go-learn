package subject_test

import (
	"task.local/replace-byte-001/subject"
	"testing"
)

func TestReplaceByteMiddle(t *testing.T) {
	input := "abc"
	got, ok := subject.ReplaceByte(input, 1, 'X')
	if !ok || got != "aXc" {
		t.Fatalf("ReplaceByte = (%q, %v), want (aXc, true)", got, ok)
	}
	if input != "abc" {
		t.Fatal("original string changed")
	}
}

func TestReplaceByteOpaqueValue(t *testing.T) {
	got, ok := subject.ReplaceByte("\x00\xfe", 1, 255)
	if !ok || got != "\x00\xff" {
		t.Fatalf("byte replacement = (%q, %v)", got, ok)
	}
}

func TestReplaceByteInvalidIndex(t *testing.T) {
	got, ok := subject.ReplaceByte("abc", -1, 'X')
	if ok || got != "abc" {
		t.Fatalf("negative index = (%q, %v), want original and false", got, ok)
	}
}
