package subject_test

import (
	"fmt"
	subject "task.local/key-summary-examples-001/subject"
)

func ExampleSummary() {
	fmt.Println(subject.Summary(map[string]int{"b": 0, "a": 4, "Z": -2}))
	// Output: Z=-2, a=4, b=0
}
func ExampleSummary_empty() {
	fmt.Println(subject.Summary(nil))
	fmt.Println(subject.Summary(map[string]int{}))
	// Output:
	// (empty)
	// (empty)
}
