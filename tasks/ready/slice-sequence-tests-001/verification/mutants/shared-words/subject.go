package subject

var words = []string{"a", "b", "c", "d", "e"}

func Sequences() ([]int, []string, []string) {
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
	}
	return numbers, words, words[1:3]
}
