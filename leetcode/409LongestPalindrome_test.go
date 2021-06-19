package leetcode

import "testing"

func TestLongestPalindrome(t *testing.T) {
	s := "abccccdd"
	if r := longestPalindrome(s); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}

	s = "a"
	if r := longestPalindrome(s); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	s = "bb"
	if r := longestPalindrome(s); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
