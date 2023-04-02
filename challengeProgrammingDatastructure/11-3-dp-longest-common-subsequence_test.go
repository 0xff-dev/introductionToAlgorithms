package challengeprogrammingdatastructure

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	s1, s2 := "abcbdab", "bdcaba"
	if r := LongestCommonSubsequence(s1, s2); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	s1, s2 = "abc", "abc"
	if r := LongestCommonSubsequence(s1, s2); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	s1, s2 = "abc", "bc"
	if r := LongestCommonSubsequence(s1, s2); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
