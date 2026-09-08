package subject_test

import (
	"task.local/receiver-speakers-001/subject"
	"testing"
)

type customSpeaker struct {
	text  string
	calls int
}

func (s *customSpeaker) Speak() string { s.calls++; return s.text }

func TestPublicCase01(t *testing.T) {
	var a subject.Speaker = subject.English{}
	var b subject.Speaker = &subject.English{}
	var c subject.Speaker = &subject.Chinese{}
	for _, x := range []struct {
		s    subject.Speaker
		want string
	}{{a, "Hello World"}, {b, "Hello World"}, {c, "你好世界"}} {
		if got := x.s.Speak(); got != x.want {
			t.Errorf("speaker got%q want%q", got, x.want)
		}
	}
	if got := subject.BuiltinGreetings(); got != ([2]string{"Hello World", "你好世界"}) {
		t.Errorf("builtins got%v", got)
	}
}

func TestPublicCase02(t *testing.T) {
	s := &customSpeaker{text: " exact "}
	if got := subject.SayHello(s); got != " exact " || s.calls != 1 {
		t.Errorf("dispatch got%q calls%d", got, s.calls)
	}
	s.text = ""
	if got := subject.SayHello(s); got != "" || s.calls != 2 {
		t.Errorf("current dispatch got%q calls%d", got, s.calls)
	}
}
