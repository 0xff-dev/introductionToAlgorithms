package leetcode

import "testing"

func TestFindTargetSumWays(t *testing.T) {
	nums, target := []int{1, 1, 1, 1, 1}, 3
	if r := findTargetSumWays(nums, target); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}

	nums, target = []int{1}, 1
	if r := findTargetSumWays(nums, target); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
