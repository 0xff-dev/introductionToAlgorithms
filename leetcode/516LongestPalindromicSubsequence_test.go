package leetcode

import "testing"

func TestLongestPalinfromicSubseq(t *testing.T) {
	s := "bbbab"
	if r := longestPalindromeSubseq(s); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	s = "cbbd"
	if r := longestPalindromeSubseq(s); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	s = "a"
	if r := longestPalindromeSubseq(s); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
