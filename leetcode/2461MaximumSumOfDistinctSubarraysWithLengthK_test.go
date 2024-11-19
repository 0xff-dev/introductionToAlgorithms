package leetcode

import "testing"

func TestMaximumSubarraySum(t *testing.T) {
	nums, k, exp := []int{1, 5, 4, 2, 9, 9, 9}, 3, int64(15)
	if r := maximumSubarraySum(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	nums, k, exp = []int{4, 4, 4}, 3, 0
	if r := maximumSubarraySum(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
