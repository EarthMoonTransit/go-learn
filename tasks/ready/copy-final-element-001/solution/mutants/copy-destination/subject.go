package subject

func CopyLast(source, destination []string) {
	destination = append([]string(nil), destination...)
	destination[len(destination)-1] = source[len(source)-1]
}
