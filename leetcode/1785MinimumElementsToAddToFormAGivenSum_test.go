package leetcode

import "testing"

func TestMinElements(t *testing.T) {
	nums := []int{1, -1, 1}
	limit, goal := 3, -4
	if r := minElements(nums, limit, goal); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	nums = []int{1, -10, 9, 1}
	limit, goal = 100, 0
	if r := minElements(nums, limit, goal); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
