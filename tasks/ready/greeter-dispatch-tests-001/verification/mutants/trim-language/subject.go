package subject

import (
	"strings"
)

type Greeter interface {
	LanguageName() string
	Greet(string) string
}

func SayHello(name string, g Greeter) string {
	return "I can speak " + strings.TrimSpace(g.LanguageName()) + ": " + g.Greet(name)
}
