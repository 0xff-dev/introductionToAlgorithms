package leetcode

import "testing"

func TestFindScore(t *testing.T) {
	nums, exp := []int{2, 1, 3, 4, 5, 2}, int64(7)
	if r := findScore(nums); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	nums, exp = []int{2, 3, 5, 1, 3, 2}, 5
	if r := findScore(nums); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
