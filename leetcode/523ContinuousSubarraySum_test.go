package leetcode

import "testing"

func TestCheckSubarraySum(t *testing.T) {
	nums := []int{23, 2, 4, 6, 7}
	if !checkSubarraySum(nums, 6) {
		t.Fatalf("expect true get false")
	}

	nums = []int{23, 2, 6, 4, 7}
	if !checkSubarraySum(nums, 6) {
		t.Fatalf("expect true get false")
	}

	nums = []int{23, 2, 6, 4, 7}
	if checkSubarraySum(nums, 13) {
		t.Fatalf("expect false get true")
	}
}
