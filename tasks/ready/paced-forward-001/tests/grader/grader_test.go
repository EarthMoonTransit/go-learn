package subject_test

import (
	"testing"
	"testing/synctest"
)

func TestPaceFiniteReleaseExtensions(t *testing.T) {
	for length := 0; length <= 8; length++ {
		synctest.Test(t, func(t *testing.T) {
			values := make([]int, length)
			for i := range values {
				values[i] = (i*5+length)%7 - 3
			}
			p := beginPace(t, values)
			defer p.finish(t, values)
			p.endRelease()
			for _, value := range values {
				p.take(t, value, "finite release closure")
			}
		})
	}
}
func TestPaceReleaseBeforeAnyInput(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		p := beginPace(t, nil)
		defer p.finish(t, []int{0, -8, 0, 8})
		p.endRelease()
		for _, value := range []int{0, -8, 0, 8} {
			p.input <- value
			p.take(t, value, "already-closed release")
		}
	})
}
