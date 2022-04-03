package leetcode

import "testing"

func TestSearchInsert(t *testing.T) {
	nums, target := []int{1, 3, 5, 6}, 5
	if r := searchInsert(nums, target); r != 2 {
		t.Fatalf("expect %d get %d", 2, r)
	}

	target = 2
	if r := searchInsert(nums, target); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	target = 7
	if r := searchInsert(nums, target); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	target = 0
	if r := searchInsert(nums, target); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}
