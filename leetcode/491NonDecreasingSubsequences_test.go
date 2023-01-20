package leetcode

import "testing"

func TestFindSubsequences(t *testing.T) {
	n := []int{4, 6, 7, 7}
	r := findSubsequences(n)
	t.Logf("%v", r)
}
