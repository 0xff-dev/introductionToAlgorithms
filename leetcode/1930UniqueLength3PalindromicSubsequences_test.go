package leetcode

import "testing"

func TestCountPalindromicSubsequence(t *testing.T) {
	s := "aabca"
	if r := countPalindromicSubsequence(s); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	s = "abc"
	if r := countPalindromicSubsequence(s); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
	s = "bbcbaba"
	if r := countPalindromicSubsequence(s); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
}
