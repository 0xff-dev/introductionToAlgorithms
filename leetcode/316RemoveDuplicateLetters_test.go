package leetcode

import "testing"

func TestRemoveDuplicateLetters(t *testing.T) {
	s := "bcabc"
	if r := removeDuplicateLetters(s); r != "abc" {
		t.Fatalf("expect '%s' get '%s'", "abc", r)
	}
	s = "cbacdcbc"
	if r := removeDuplicateLetters(s); r != "acdb" {
		t.Fatalf("expect acdb get %s", r)
	}
}
