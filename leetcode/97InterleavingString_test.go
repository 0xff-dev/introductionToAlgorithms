package leetcode

import "testing"

func TestIsInterleave(t *testing.T) {
	s1, s2, s3 := "aabcc", "dbbca", "aadbbcbcac"
	if !isInterleave(s1, s2, s3) {
		t.Fatalf("s1: %s, s2: %s, s3: %s. expect true get false", s1, s2, s3)
	}
	s1, s2, s3 = "aabcc", "dbbca", "aadbbbaccc"
	if isInterleave(s1, s2, s3) {
		t.Fatalf("s1: %s, s2: %s, s3: %s. expect false get true", s1, s2, s3)
	}
	s1, s2, s3 = "a", "", "a"
	if !isInterleave(s1, s2, s3) {
		t.Fatalf("s1: %s, s2: %s, s3: %s. expect true get false", s1, s2, s3)
	}
	s1, s2, s3 = "aa", "ab", "abaa"
	if !isInterleave(s1, s2, s3) {
		t.Fatalf("s1: %s, s2: %s, s3: %s. expect true get false", s1, s2, s3)
	}
	s1, s2, s3 = "ab", "bc", "babc"
	if !isInterleave(s1, s2, s3) {
		t.Fatalf("s1: %s, s2: %s, s3: %s. expect true get false", s1, s2, s3)
	}
}
