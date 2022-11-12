package leetcode

import "testing"

func TestIsValid1003(t *testing.T) {
	s := "aabcbc"
	if !isValid1003(s) {
		t.Fatalf("expect true get false")
	}

	s = "abcabcababcc"
	if !isValid1003(s) {
		t.Fatalf("expect true get false")
	}

	s = "abccba"
	if isValid1003(s) {
		t.Fatalf("expect false get true")
	}

	s = "ab"
	if isValid1003(s) {
		t.Fatalf("expect false get true")
	}

	s = "a"
	if isValid1003(s) {
		t.Fatalf("expect false get true")
	}
}
