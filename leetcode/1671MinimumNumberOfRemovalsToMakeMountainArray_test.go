package leetcode

import "testing"

func TestMinimumMountainRemovals(t *testing.T) {
	nums, exp := []int{1, 3, 1}, 0
	if r := minimumMountainRemovals(nums); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	nums, exp = []int{2, 1, 1, 5, 6, 2, 3, 1}, 3
	if r := minimumMountainRemovals(nums); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
