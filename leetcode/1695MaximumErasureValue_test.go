package leetcode

import "testing"

func TestMaximumUniqueSubarray(t *testing.T) {
	nums := []int{4, 2, 4, 5, 6}
	if r := maximumUniqueSubarray(nums); r != 17 {
		t.Fatalf("expect 17 get %d", r)
	}

	nums = []int{5, 2, 1, 2, 5, 2, 1, 2, 5}
	if r := maximumUniqueSubarray(nums); r != 8 {
		t.Fatalf("expect 8 get %d", r)
	}

	nums = []int{1}
	if r := maximumUniqueSubarray(nums); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	nums = []int{1, 2, 1}
	if r := maximumUniqueSubarray(nums); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	nums = []int{1, 1, 1, 1, 1, 3, 1, 1, 3, 3, 4}
	if r := maximumUniqueSubarray(nums); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}
}
