package leetcode

import "testing"

func TestFindLongestWord(t *testing.T) {
	s, d := "abpcplea", []string{"alp", "apple", "monkey", "plea"}
	t.Log(findLongestWord(s, d))

	s, d = "abpcplea", []string{"a", "b", "c"}
	t.Log(findLongestWord(s, d))

	s, d = "bab", []string{"ba", "ab", "a", "b"}
	t.Log(findLongestWord(s, d))

	s, d = "aewfafwafjlwajflwajflwafj", []string{"apple", "ewaf", "awefawfwaf", "awef", "awefe", "ewafeffewafewf"}
	t.Log(findLongestWord(s, d))
}
