package leetcode

import "testing"

func TestFindKthLargest(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	if r := findKthLargest(nums, 2); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}

	nums = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	if r := findKthLargest(nums, 4); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
}
