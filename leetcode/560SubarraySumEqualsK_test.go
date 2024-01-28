package leetcode

import "testing"

func TestSubarraySum(t *testing.T) {
	nums := []int{1, 1, 1}
	k := 2
	if r := subarraySum(nums, k); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	nums = []int{1, 2, 3}
	k = 3
	if r := subarraySum(nums, k); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
