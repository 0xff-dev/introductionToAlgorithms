package leetcode

import (
	"reflect"
	"testing"
)

func TestConstructor352(t *testing.T) {
	sr := Constructor352()
	sr.AddNum(1)
	exp := [][]int{
		{1, 1},
	}
	if r := sr.GetIntervals(); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	sr.AddNum(3)
	exp = [][]int{
		{1, 1}, {3, 3},
	}
	if r := sr.GetIntervals(); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	sr.AddNum(7)
	exp = [][]int{
		{1, 1}, {3, 3}, {7, 7},
	}
	if r := sr.GetIntervals(); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	sr.AddNum(2)
	exp = [][]int{
		{1, 3}, {7, 7},
	}
	if r := sr.GetIntervals(); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	sr.AddNum(6)
	exp = [][]int{
		{1, 3}, {6, 7},
	}
	if r := sr.GetIntervals(); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
