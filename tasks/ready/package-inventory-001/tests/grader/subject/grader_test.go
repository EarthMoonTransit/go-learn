package subject_test

import "task.local/package-inventory-001/stock"

import (
	"task.local/package-inventory-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	item := stock.New("q", 99)
	item.SetSold(4)
	item.SetOnHand(6)
	if got := subject.Describe(item); got != "q:99:6:4" {
		t.Errorf("setter order got%q", got)
	}
}
