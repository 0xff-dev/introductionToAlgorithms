package leetcode

import "testing"

func TestLongestPalindrome(t *testing.T) {
	for _, tc := range []string{"babad", "cbbd", "a", "ac"} {
		t.Logf("input[%s]=[%s]", tc, longestPalindrome(tc))
	}
}
