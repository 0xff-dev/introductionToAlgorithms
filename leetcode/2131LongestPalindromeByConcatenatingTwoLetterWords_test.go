package leetcode

import "testing"

func TestLongestPalindrome2131(t *testing.T) {
	words := []string{"dd", "aa", "bb", "dd", "aa", "dd", "bb", "dd", "aa", "cc", "bb", "cc", "dd", "cc"}
	if r := longestPalindrome2131(words); r != 22 {
		t.Fatalf("expect 20 get %d", r)
	}
}
