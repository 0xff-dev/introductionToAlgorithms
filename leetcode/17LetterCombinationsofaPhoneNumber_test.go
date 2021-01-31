package leetcode

import "testing"

func TestLetterCombinations(t *testing.T) {
	for _, testCase := range []string{
		"23",
		"",
		"2",
		"789",
	} {
		t.Log(letterCombinations(testCase))
	}
}
