package subject

func CopyLast(source, destination []string) {
	destination[len(destination)-1], source[len(source)-1] = source[len(source)-1], destination[len(destination)-1]
}
