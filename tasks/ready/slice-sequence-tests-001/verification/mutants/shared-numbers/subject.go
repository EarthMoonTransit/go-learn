package subject

var numbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func Sequences() ([]int, []string, []string) {
	words := []string{"a", "b", "c", "d", "e"}
	return numbers, words, words[1:3]
}
