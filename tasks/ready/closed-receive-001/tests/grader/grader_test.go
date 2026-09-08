package subject_test

import (
	"slices"
	"task.local/closed-receive-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	c := make(chan int, 2)
	c <- 8
	c <- 9
	close(c)
	if got := subject.CollectClosed(c); !slices.Equal(got, []int{8, 9}) {
		t.Error("send order changed")
	}
}
