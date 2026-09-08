package subject_test

import (
	"task.local/greeter-dispatch-tests-001/subject"
	"testing"
)

type greeting struct{ language, prefix string }

func (g greeting) LanguageName() string     { return g.language }
func (g greeting) Greet(name string) string { return g.prefix + name + "!" }

type echoGreeting struct{}

func (echoGreeting) LanguageName() string     { return "Echo" }
func (echoGreeting) Greet(name string) string { return name + name }

func TestContractCase01(t *testing.T) {
	got := subject.SayHello("Flora", greeting{"Italian", "Ciao "})
	want := "I can speak Italian: Ciao Flora!"
	if got != want {
		t.Errorf("Italian got %q want %q", got, want)
	}
}

func TestContractCase02(t *testing.T) {
	for _, c := range []struct {
		n    string
		g    subject.Greeter
		want string
	}{{"Fabrício", greeting{"Portuguese", "Olá "}, "I can speak Portuguese: Olá Fabrício!"}, {" a ", echoGreeting{}, "I can speak Echo:  a  a "}, {"", greeting{"", ""}, "I can speak : !"}, {" x ", greeting{" L ", "["}, "I can speak  L : [ x !"}} {
		if got := subject.SayHello(c.n, c.g); got != c.want {
			t.Errorf("dispatch got %q want %q", got, c.want)
		}
	}
}
