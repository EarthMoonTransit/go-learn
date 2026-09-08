package subject_test

import (
	"slices"
	"task.local/text-level-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	for n := 0; n <= 9; n++ {
		value := subject.Level(n)
		data, e := value.MarshalText()
		if e != nil || len(data) != 1 || data[0] != byte(n)+'0' {
			t.Errorf("marshal%d got%q,%v", n, data, e)
		}
		var got subject.Level
		if e = got.UnmarshalText([]byte{byte(n) + '0'}); e != nil || got != value {
			t.Errorf("unmarshal%d got%d,%v", n, got, e)
		}
	}
}

func TestPublicCase02(t *testing.T) {
	sentinel := originalSentinel
	if sentinel == nil {
		t.Fatal("nil level sentinel")
	}
	for _, text := range []string{"", "10", " 1", "1 ", "+1", "１", "/", ":"} {
		l := subject.Level(7)
		input := []byte(text)
		before := slices.Clone(input)
		e := l.UnmarshalText(input)
		if e != sentinel || l != 7 || !slices.Equal(input, before) || subject.ErrLevel != sentinel {
			t.Errorf("invalid %q got%d,%v", text, l, e)
		}
	}
	for _, l := range []subject.Level{10, 255} {
		data, e := l.MarshalText()
		if data != nil || e != sentinel || subject.ErrLevel != sentinel {
			t.Errorf("invalid level%d got%q,%v", l, data, e)
		}
	}
}

func TestPublicCase03(t *testing.T) {
	a, e := subject.Level(1).MarshalText()
	b, f := subject.Level(1).MarshalText()
	if e != nil || f != nil || len(a) != 1 || len(b) != 1 {
		t.Fatal("bad marshal")
	}
	a[0] = '9'
	if b[0] != '1' {
		t.Error("marshal outputs share storage")
	}
	input := []byte("2")
	var l subject.Level
	if l.UnmarshalText(input) != nil || string(input) != "2" || l != 2 {
		t.Errorf("valid input changed %q level%d", input, l)
	}
}

var originalSentinel = subject.ErrLevel
