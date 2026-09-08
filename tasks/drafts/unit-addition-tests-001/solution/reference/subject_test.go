package subject_test

import (
	"task.local/unit-addition-tests-001/subject"
	"testing"
)

func TestContractCase01(t *testing.T) {
	v := map[string]int{}
	u := map[string]int{"dozen": 12}
	if !subject.AddUnit(v, u, "item", "dozen") || v["item"] != 12 {
		t.Errorf("add got %v want item:12 and true", v)
	}
}

func TestContractCase02(t *testing.T) {
	v := map[string]int{"a": 4, "b": 3}
	u := map[string]int{"triple": 3}
	if !subject.AddUnit(v, u, "a", "triple") || !subject.AddUnit(v, u, "a", "triple") || v["a"] != 10 || v["b"] != 3 {
		t.Errorf("accumulation got %v want a:10 b:3", v)
	}
}

func TestContractCase03(t *testing.T) {
	v := map[string]int{"a": 4}
	if subject.AddUnit(v, map[string]int{}, "a", "missing") || len(v) != 1 || v["a"] != 4 {
		t.Errorf("unknown unit changed %v or returned true", v)
	}
}

func TestContractCase04(t *testing.T) {
	v := map[string]int{}
	if subject.AddUnit(v, nil, "a", "missing") || len(v) != 0 {
		t.Errorf("nil units changed %v or returned true", v)
	}
}

func TestContractCase05(t *testing.T) {
	v := map[string]int{}
	ok := subject.AddUnit(v, map[string]int{"zero": 0}, "a", "zero")
	n, present := v["a"]
	if !ok || !present || n != 0 {
		t.Errorf("zero unit got %v,%v want present a:0 and true", v, ok)
	}
}

func TestContractCase06(t *testing.T) {
	v := map[string]int{}
	u := map[string]int{"triple": 3, "other": 7}
	subject.AddUnit(v, u, "a", "triple")
	if len(u) != 2 || u["triple"] != 3 || u["other"] != 7 {
		t.Errorf("units changed: %v", u)
	}
}
