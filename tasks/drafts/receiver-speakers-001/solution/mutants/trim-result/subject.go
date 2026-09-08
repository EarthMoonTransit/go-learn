package subject

import (
	"strings"
)

type Speaker interface{ Speak() string }
type English struct{}
type Chinese struct{}

func (English) Speak() string     { return "Hello World" }
func (*Chinese) Speak() string    { return "你好世界" }
func SayHello(s Speaker) string   { return strings.TrimSpace(s.Speak()) }
func BuiltinGreetings() [2]string { return [2]string{SayHello(English{}), SayHello(&Chinese{})} }
