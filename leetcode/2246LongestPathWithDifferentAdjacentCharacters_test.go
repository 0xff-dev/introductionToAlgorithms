package leetcode

import "testing"

func TestLongestPath(t *testing.T) {
	parent, s := []int{-1, 0, 0, 1, 1, 2}, "abacbe"
	if r := longestPath(parent, s); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	parent, s = []int{-1, 0, 0, 0}, "aabc"
	if r := longestPath(parent, s); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}
