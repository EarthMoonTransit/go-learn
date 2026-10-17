package subject_test

import (
	"task.local/replace-pending-value-001/subject"
	"testing"
	"testing/synctest"
)

func checkReplacement(t *testing.T, old, next int) {
	t.Helper()
	synctest.Test(t, func(t *testing.T) {
		ch := make(chan int, 1)
		ch <- old
		defer func() {
			// Every required defective call uses at most one blocking send.
			for step := 0; step < 3; step++ {
				select {
				case _, open := <-ch:
					if !open {
						return
					}
				default:
				}
				synctest.Wait()
			}
		}()
		returned := make(chan int, 1)
		go func() { returned <- subject.ReplacePending(ch, next) }()
		synctest.Wait()
		var got int
		select {
		case got = <-returned:
		default:
			t.Errorf("ReplacePending(%d,%d) blocked on full buffer, want receive before send", old, next)
			select {
			case <-ch:
			default:
			}
			synctest.Wait()
			select {
			case got = <-returned:
			default:
				t.Fatal("ReplacePending remains blocked after cleanup drain, want completed call")
			}
		}
		if got != old {
			t.Errorf("ReplacePending(%d,%d) returned %d, want old %d", old, next, got, old)
		}
		select {
		case value, open := <-ch:
			if !open {
				t.Error("replacement channel closed, want open with next pending")
			} else if value != next {
				t.Errorf("pending value = %d, want next %d", value, next)
			}
		default:
			t.Errorf("pending value = absent, want next %d", next)
		}
		select {
		case value, open := <-ch:
			if !open {
				t.Error("channel after drain = closed, want open")
			} else {
				t.Errorf("extra pending value = %d, want none", value)
			}
		default:
		}
	})
}
func TestPendingReplacement(t *testing.T) {
	for _, c := range [][2]int{{0, 7}, {-3, 8}, {5, -2}, {4, 4}} {
		checkReplacement(t, c[0], c[1])
	}
}
