package subject_test

import (
	"task.local/text-border-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	if got := subject.AddBorder("Welcome!", 10); got != "**********\nWelcome!\n**********" {
		t.Errorf("result got %q want %q", got, "**********\nWelcome!\n**********")
	}
}

func TestPublicCase02(t *testing.T) {
	if got := subject.AddBorder("Hi", 2); got != "**\nHi\n**" {
		t.Errorf("result got %q want %q", got, "**\nHi\n**")
	}
}

func TestPublicCase03(t *testing.T) {
	if got := subject.AddBorder("x", 0); got != "\nx\n" {
		t.Errorf("result got %q want %q", got, "\nx\n")
	}
}

func TestPublicCase04(t *testing.T) {
	if got := subject.AddBorder("", 1); got != "*\n\n*" {
		t.Errorf("result got %q want %q", got, "*\n\n*")
	}
}

func TestPublicCase05(t *testing.T) {
	if got := subject.AddBorder(" a\nb ", 2); got != "**\n a\nb \n**" {
		t.Errorf("result got %q want %q", got, "**\n a\nb \n**")
	}
}
