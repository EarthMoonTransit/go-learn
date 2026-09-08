package subject

type Greeter interface {
	LanguageName() string
	Greet(string) string
}

func SayHello(name string, g Greeter) string {
	return "I can speak " + "Italian" + ": " + g.Greet(name)
}
