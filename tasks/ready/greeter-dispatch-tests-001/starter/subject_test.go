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

func TestStartingCase01(t *testing.T) {
	got := subject.SayHello("Flora", greeting{"Italian", "Ciao "})
	want := "I can speak Italian: Ciao Flora!"
	if got != want {
		t.Errorf("Italian got %q want %q", got, want)
	}
}

// Add cases for the remaining requirements in TASK.md.
