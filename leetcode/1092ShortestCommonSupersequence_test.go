package leetcode

import "testing"

func TestSortestCommonSupersequence(t *testing.T) {
	s1, s2 := "abac", "cab"
	r := shortestCommonSupersequence(s1, s2)
	t.Logf("%s\n", r)
	s1, s2 = "a", "b"
	r = shortestCommonSupersequence(s1, s2)
	t.Logf("%s\n", r)
	s1, s2 = "a", "a"
	t.Logf("%s\n", shortestCommonSupersequence(s1, s2))
}
