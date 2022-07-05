package leetcode

import "testing"

func TestLongestConsecutive(t *testing.T) {
	nums := []int{100, 4, 200, 1, 3, 2}
	if r := longestConsecutive(nums); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	nums = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	if r := longestConsecutive(nums); r != 9 {
		t.Fatalf("expect 9 get %d", r)
	}

	nums = []int{1, 2, 0, 1}
	if r := longestConsecutive(nums); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}
