package leetcode

import "testing"

type testCase struct {
	Input string
	Ans   int
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, tc := range []testCase{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"abcdefg", 7},
		{"cdd", 2},
		{"umvejcuuk", 6},
		{" ", 1},
		{"abba", 2},
	} {
		if ans := lengthOfLongestSubstring1(tc.Input); ans != tc.Ans {
			t.Fatalf("input[%s] expect: %d, get: %d", tc.Input, tc.Ans, ans)
		}
	}
}
