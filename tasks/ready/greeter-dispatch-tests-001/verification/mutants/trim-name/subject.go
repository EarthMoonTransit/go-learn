package subject

import (
	"strings"
)

type Greeter interface {
	LanguageName() string
	Greet(string) string
}

func SayHello(name string, g Greeter) string {
	return "I can speak " + g.LanguageName() + ": " + g.Greet(strings.TrimSpace(name))
}
