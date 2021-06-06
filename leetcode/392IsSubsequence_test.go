package leetcode

import "testing"

func TestIsSubsequence(t *testing.T) {
	s, t1 := "abc", "ahbgdc"
	if !isSubsequence(s, t1) {
		t.Fatalf("expect true get false")
	}

	s, t1 = "axc", "ahbgdc"
	if isSubsequence(s, t1) {
		t.Fatalf("expect false get true")
	}

	s, t1 = "abc", ""
	if isSubsequence(s, t1) {
		t.Fatalf("expect false get true")
	}
	s, t1 = "", "abc"
	if !isSubsequence(s, t1) {
		t.Fatalf("expect true get false")
	}
}
