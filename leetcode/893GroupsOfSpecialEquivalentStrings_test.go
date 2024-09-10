package leetcode

import "testing"

func TestNumSpecialEquivGroups(t *testing.T) {
	words, exp := []string{"abcd", "cdab", "cbad", "xyzz", "zzxy", "zzyx"}, 3
	if r := numSpecialEquivGroups(words); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	words, exp = []string{"abc", "acb", "bac", "bca", "cab", "cba"}, 3
	if r := numSpecialEquivGroups(words); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
