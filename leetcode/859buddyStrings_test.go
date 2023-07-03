package leetcode

import "testing"

func TestBuddyStrings(t *testing.T) {
	s, goal := "ab", "ba"
	if !buddyStrings(s, goal) {
		t.Fatalf("expect true get false")
	}

	s, goal = "ab", "ab"
	if buddyStrings(s, goal) {
		t.Fatalf("expect false get true")
	}
	s, goal = "aa", "aa"
	if !buddyStrings(s, goal) {
		t.Fatalf("expect true get false")
	}
}
