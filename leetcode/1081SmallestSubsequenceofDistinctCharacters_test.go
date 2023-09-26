package leetcode

import "testing"

func TestSmallestSubsequence(t *testing.T) {
	s := "bcabc"
	if r := smallestSubsequence(s); r != "abc" {
		t.Fatalf("expect abc get %s", r)
	}

	s = "cbacdcbc"
	if r := smallestSubsequence(s); r != "acdb" {
		t.Fatalf("expect acdb get %s", r)
	}
}
