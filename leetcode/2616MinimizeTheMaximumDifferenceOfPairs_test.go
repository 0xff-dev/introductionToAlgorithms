package leetcode

import "testing"

func TestMinimizeMax(t *testing.T) {
	nums := []int{10, 1, 2, 7, 1, 3}
	p := 2
	if r := minimizeMax(nums, p); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	nums = []int{4, 2, 1, 2}
	p = 1
	if r := minimizeMax(nums, p); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}
