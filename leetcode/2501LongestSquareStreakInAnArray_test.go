package leetcode

import "testing"

func TestLongestSquareStreak(t *testing.T) {
	nums, exp := []int{4, 3, 6, 16, 8, 2}, 3
	if r := longestSquareStreak(nums); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	nums, exp = []int{2, 3, 5, 6, 7}, -1
	if r := longestSquareStreak(nums); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
