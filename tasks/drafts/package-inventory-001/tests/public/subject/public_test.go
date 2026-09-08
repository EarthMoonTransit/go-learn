package subject_test

import "task.local/package-inventory-001/stock"

import (
	"task.local/package-inventory-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	item := stock.New("part", 4)
	if item.Name != "part" || item.Weight != 4 || item.OnHand() != 0 || item.Sold() != 0 {
		t.Errorf("new got%+v counts%d,%d", item, item.OnHand(), item.Sold())
	}
	item.SetOnHand(8)
	item.SetSold(3)
	if got := subject.Describe(item); got != "part:4:8:3" {
		t.Errorf("summary got%q wantpart:4:8:3", got)
	}
}

func TestPublicCase02(t *testing.T) {
	item := stock.New(" x ", 0)
	other := stock.New("", 2)
	item.SetOnHand(9)
	item.SetSold(2)
	item.SetSold(0)
	if item.OnHand() != 9 || item.Sold() != 0 || item.Name != " x " || item.Weight != 0 {
		t.Errorf("set sold changed other state %+v", item)
	}
	item.SetOnHand(0)
	other.SetSold(7)
	if got := subject.Describe(item); got != " x :0:0:0" {
		t.Errorf("first got%q", got)
	}
	if got := subject.Describe(other); got != ":2:0:7" {
		t.Errorf("independent got%q", got)
	}
}
