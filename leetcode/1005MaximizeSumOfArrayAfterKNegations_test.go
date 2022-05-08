package leetcode

import "testing"

func TestLargestSumAfterKNegation(t *testing.T) {
	nums := []int{4, 2, 3}
	if r := largestSumAfterKNegations(nums, 1); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}

	nums = []int{3, -1, 0, 2}
	if r := largestSumAfterKNegations(nums, 3); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}

	nums = []int{2, -3, -1, 5, -4}
	if r := largestSumAfterKNegations(nums, 2); r != 13 {
		t.Fatalf("expect 13 get %d", r)
	}
}
