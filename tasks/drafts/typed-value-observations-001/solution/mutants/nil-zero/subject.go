package subject

func ZeroValues() (any, any, any) {
	var text string
	var number int
	var flag bool
	_ = text
	_ = number
	_ = flag
	return nil, nil, nil
}
func LiteralValues() (any, any, any, float32) {
	text := "go"
	number := 42
	flag := true
	small := float32(3.14)
	return text, number, flag, small
}
