package leetcode

import "testing"

func TestMinCost2448(t *testing.T) {
	nums, cost := []int{1, 3, 5, 2}, []int{2, 3, 1, 14}
	if r := minCost2448(nums, cost); r != 8 {
		t.Fatalf("expect 8 get %d", r)
	}
}
