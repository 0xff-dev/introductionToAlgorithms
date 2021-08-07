package leetcode

import "testing"

func TestShortestPalindrome(t *testing.T) {
	s := "aacecaaa"
	if r := shortestPalindrome(s); r != "aaacecaaa" {
		t.Fatalf("expect aaacecaaa get %s", r)
	}

	s = "abcd"
	if r := shortestPalindrome(s); r != "dcbabcd" {
		t.Fatalf("expect dcbabcd get %s", r)
	}
}
