package subject_test

import (
	"fmt"
	subject "task.local/key-summary-examples-001/subject"
)

func ExampleSummary_empty() { fmt.Println(subject.Summary(nil)) /* Output: (empty) */ }
