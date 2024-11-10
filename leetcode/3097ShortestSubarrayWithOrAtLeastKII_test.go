package leetcode

import "testing"

func TestMinimumSubarrayLength(t *testing.T) {
	nums, k, exp := []int{1, 2, 3}, 2, 1
	if r := minimumSubarrayLength(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	nums, k, exp = []int{2, 1, 8}, 10, 3
	if r := minimumSubarrayLength(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	nums, k, exp = []int{1, 2}, 0, 1
	if r := minimumSubarrayLength(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
