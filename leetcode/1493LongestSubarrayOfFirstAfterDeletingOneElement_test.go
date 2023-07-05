package leetcode

import "testing"

func TestLongestSubarray(t *testing.T) {
	nums := []int{1, 1, 0, 1}
	if r := longestSubarray(nums); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	nums = []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	if r := longestSubarray(nums); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}

	nums = []int{1, 1, 1}
	if r := longestSubarray(nums); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
