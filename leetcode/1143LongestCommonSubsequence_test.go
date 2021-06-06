package leetcode

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	t1, t2 := "abcde", "ace"
	if r := longestCommonSubsequence(t1, t2); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	t1, t2 = "abc", "abc"
	if r := longestCommonSubsequence(t1, t2); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	t1, t2 = "abc", "dev"
	if r := longestCommonSubsequence(t1, t2); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}
