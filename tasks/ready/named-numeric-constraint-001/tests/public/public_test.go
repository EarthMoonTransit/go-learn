package subject_test

import (
	"task.local/named-numeric-constraint-001/subject"
	"testing"
)

func TestPublicNumericValues(t *testing.T) {
	if subject.Double(0) != 0 || subject.Double(-3) != -6 || subject.Double(1.25) != 2.5 {
		t.Error("numeric doubling incorrect")
	}
}
