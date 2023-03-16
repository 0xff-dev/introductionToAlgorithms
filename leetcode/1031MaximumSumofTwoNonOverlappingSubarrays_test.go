package leetcode

import "testing"

func TestMaxSumTwoNoOverlap(t *testing.T) {
	nums := []int{0, 6, 5, 2, 2, 5, 1, 9, 4}
	f, s := 1, 2
	if r := maxSumTwoNoOverlap(nums, f, s); r != 20 {
		t.Fatalf("expect 20 get %d", r)
	}
	nums = []int{3, 8, 1, 3, 2, 1, 8, 9, 0}
	f, s = 3, 2
	if r := maxSumTwoNoOverlap(nums, f, s); r != 29 {
		t.Fatalf("expect 29 get %d", r)
	}

	nums = []int{2, 1, 5, 6, 0, 9, 5, 0, 3, 8}
	f, s = 4, 3
	if r := maxSumTwoNoOverlap(nums, f, s); r != 31 {
		t.Fatalf("expect 31 get %d", r)
	}
}
