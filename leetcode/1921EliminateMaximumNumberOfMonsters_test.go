package leetcode

import "testing"

func TestEliminateMaximum(t *testing.T) {
	d, s := []int{1, 3, 4}, []int{1, 1, 1}
	if r := eliminateMaximum(d, s); r != 3 {
		t.Fatalf("expect %d get %d", 3, r)
	}
	d, s = []int{1, 1, 2, 3}, []int{1, 1, 1, 1}
	if r := eliminateMaximum(d, s); r != 1 {
		t.Fatalf("expect %d get %d", 1, r)
	}
	d, s = []int{3, 2, 4}, []int{5, 3, 2}
	if r := eliminateMaximum(d, s); r != 1 {
		t.Fatalf("expect %d get %d", 1, r)
	}
}
