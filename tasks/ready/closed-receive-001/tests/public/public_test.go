package subject_test

import (
	"slices"
	"task.local/closed-receive-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	c := make(chan int, 4)
	for _, v := range []int{3, 0, -2, 7} {
		c <- v
	}
	close(c)
	if got := subject.CollectClosed(c); !slices.Equal(got, []int{3, 0, -2, 7}) {
		t.Errorf("drain got%v want[3 0 -2 7]", got)
	}
	if got := subject.CollectClosed(c); len(got) != 0 {
		t.Errorf("already drained got%v", got)
	}
}

func TestPublicCase02(t *testing.T) {
	c := make(chan int)
	close(c)
	if got := subject.CollectClosed(c); len(got) != 0 {
		t.Errorf("closed empty got%v", got)
	}
	one := make(chan int, 1)
	one <- 0
	close(one)
	if got := subject.CollectClosed(one); !slices.Equal(got, []int{0}) {
		t.Errorf("queued zero got%v", got)
	}
}
