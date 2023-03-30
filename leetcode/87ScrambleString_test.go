package leetcode

import "testing"

func TestIsScramble(t *testing.T) {
	s1, s2 := "great", "rgeat"
	if !isScramble(s1, s2) {
		t.Fatalf("expect true get false")
	}

	s1, s2 = "abcde", "cadbd"
	if isScramble(s1, s2) {
		t.Fatalf("expect false get true")
	}

	s1, s2 = "a", "a"
	if !isScramble(s1, s2) {
		t.Fatalf("expect true get false")
	}
}
