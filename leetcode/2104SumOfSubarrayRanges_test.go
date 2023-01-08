package leetcode

import "testing"

func TestSubArrayRanges(t *testing.T) {
	nums := []int{1, 2, 3}
	if r := subArrayRanges(nums); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	nums = []int{1, 3, 3}
	if r := subArrayRanges(nums); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	nums = []int{4, -2, -3, 4, 1}
	if r := subArrayRanges(nums); r != 59 {
		t.Fatalf("expect 59  get %d", r)
	}
}
