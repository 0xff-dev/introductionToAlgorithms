package leetcode

import "testing"

func TestLongestOnes(t *testing.T) {
	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	k := 2
	if r := longestOnes(nums, k); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}
	nums = []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	k = 3
	if r := longestOnes(nums, k); r != 10 {
		t.Fatalf("expect 10 get %d", r)
	}
}
