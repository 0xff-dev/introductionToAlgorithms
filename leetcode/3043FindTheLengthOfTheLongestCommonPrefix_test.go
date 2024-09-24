package leetcode

import "testing"

func TestLongestCommonPrefix3043(t *testing.T) {
	a1, a2, exp := []int{1, 10, 100}, []int{1000}, 3
	if r := longestCommonPrefix3043(a1, a2); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	a1, a2, exp = []int{1, 2, 3}, []int{4, 4, 4}, 0
	if r := longestCommonPrefix3043(a1, a2); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
