package leetcode

import "testing"

func TestCountMaxOrSubsets(t *testing.T) {
	nums := []int{3, 1}
	if r := countMaxOrSubsets(nums); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	nums = []int{2, 2, 2}
	if r := countMaxOrSubsets(nums); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}
	nums = []int{3, 1, 2, 5}
	if r := countMaxOrSubsets(nums); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}
}
