package leetcode

import "testing"

func TestCloseStrings(t *testing.T) {
	s1, s2 := "abc", "bca"
	if !closeStrings(s1, s2) {
		t.Fatalf("expect true get false")
	}

	s1, s2 = "a", "aa"
	if closeStrings(s1, s2) {
		t.Fatalf("expect false get true")
	}

	s1, s2 = "cabbba", "abbccc"
	if !closeStrings(s1, s2) {
		t.Fatalf("expect true get flase")
	}
}
