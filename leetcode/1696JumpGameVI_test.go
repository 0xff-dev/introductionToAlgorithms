package leetcode

import "testing"

func TestMaxResult(t *testing.T) {
	nums := []int{1, -1, -2, 4, -7, 3}
	if r := maxResult(nums, 2); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}

	nums = []int{10, -5, -2, 4, 0, 3}
	if r := maxResult(nums, 3); r != 17 {
		t.Fatalf("expect 17 get %d", r)
	}

	nums = []int{1, -5, -20, 4, -1, 3, -6, -3}
	if r := maxResult(nums, 2); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}
