package subject_test

import (
	"fmt"
	"task.local/nested-stringers-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := fmt.Sprint(subject.Temperature{Degree: 75, Unit: subject.Fahrenheit}); got != "75 °F" {
		t.Errorf("fahrenheit got %q want75 °F", got)
	}
}
