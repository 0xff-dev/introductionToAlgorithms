package leetcode

import "testing"

func TestSumSubarrayMins(t *testing.T) {
	nums := []int{3, 4, 5, 2, 6, 1}
	if r := sumSubarrayMins(nums); r != 50 {
		t.Fatalf("expect 50 get %d", r)
	}

	nums = []int{3, 1, 2, 4}
	if r := sumSubarrayMins(nums); r != 17 {
		t.Fatalf("expect 17 get %d", r)
	}
}
