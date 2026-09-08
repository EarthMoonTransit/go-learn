package subject_test

import (
	"task.local/month-enumeration-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	values := []int{subject.Jan, subject.Feb, subject.Mar, subject.Apr, subject.May, subject.Jun, subject.Jul, subject.Aug, subject.Sep, subject.Oct, subject.Nov, subject.Dec}
	for i, v := range values {
		if v != i+1 || subject.GetMonth(v) != i+1 {
			t.Errorf("month%d constant%d result%d", i+1, v, subject.GetMonth(v))
		}
	}
}
