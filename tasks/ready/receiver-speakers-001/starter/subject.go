package subject

type Speaker interface{ Speak() string }
type English struct{}
type Chinese struct{}

func (English) Speak() string     { return "" }
func (*Chinese) Speak() string    { return "" }
func SayHello(s Speaker) string   { return "" }
func BuiltinGreetings() [2]string { return [2]string{} }
