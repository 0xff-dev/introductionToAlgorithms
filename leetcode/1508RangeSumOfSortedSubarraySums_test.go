package leetcode

import "testing"

func TestRangeSum(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	n, left, right := 4, 1, 5
	if r := rangeSum(nums, n, left, right); r != 13 {
		t.Fatalf("expect 13 get %d", r)
	}
	nums = []int{1, 2, 3, 4}
	n, left, right = 4, 3, 4
	if r := rangeSum(nums, n, left, right); r != 6 {
		t.Fatalf("expect 13 get %d", r)
	}
	nums = []int{1, 2, 3, 4}
	n, left, right = 4, 1, 10
	if r := rangeSum(nums, n, left, right); r != 50 {
		t.Fatalf("expect 13 get %d", r)
	}
}
