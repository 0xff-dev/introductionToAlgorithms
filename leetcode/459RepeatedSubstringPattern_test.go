package leetcode

import "testing"

func TestRepeatedSubstringPattern(t *testing.T) {
	s := "abab"
	if !repeatedSubstringPattern(s) {
		t.Fatal("expect true get false")
	}

	s = "aba"
	if repeatedSubstringPattern(s) {
		t.Fatal("expect false get true")
	}

	s = "abcabcabcabc"
	if !repeatedSubstringPattern(s) {
		t.Fatal("expect true get false")
	}
}
