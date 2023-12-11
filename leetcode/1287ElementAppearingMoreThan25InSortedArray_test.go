package leetcode

import "testing"

func TestFindSpecialInteger(t *testing.T) {
	arr := []int{1, 2, 2, 6, 6, 6, 6, 7, 10}
	if r := findSpecialInteger(arr); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}
	arr = []int{1, 1}
	if r := findSpecialInteger(arr); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
