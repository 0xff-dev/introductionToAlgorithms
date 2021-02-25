package leetcode

import "testing"

func TestIsMatch(t *testing.T) {
	s, p := "aa", "*"
	if !isMatch(s, p) {
		t.Fatal("should return true")
	}
	s, p = "cb", "?a"
	if isMatch(s, p) {
		t.Fatal("should return false")
	}

	s, p = "adceb", "*a*b"
	if !isMatch(s, p) {
		t.Fatal("should return true")
	}
	s, p = "acdcb", "a*c?b"
	if isMatch(s, p) {
		t.Fatal("should return false")
	}

	s, p = "aa", "aa"
	if !isMatch(s, p) {
		t.Fatal("should return true")
	}
	s, p = "abacad", "a?a?a?"
	if !isMatch(s, p) {
		t.Fatal("should return true")
	}

	s, p = "", ""
	if !isMatch(s, p) {
		t.Fatal("should return true")
	}

	s, p = "aaaaa", "?*"
	if !isMatch(s, p) {
		t.Fatal("should return true")
	}
	s, p = "", "?"
	if isMatch(s, p) {
		t.Fatal("should return false")
	}

	s, p = "asf", ""
	if isMatch(s, p) {
		t.Fatal("should return false")
	}

	s, p = "aa", "a"
	if isMatch(s, p) {
		t.Fatal("should return false")
	}

	s, p = "", "****"
	if !isMatch(s, p) {
		t.Fatal("shoule return true")
	}

	s, p = "b", "*?*?*"
	t.Log("match: ", isMatch(s, p))
	s, p = "abcabczzzde", "*abc???de*"
	if !isMatch(s, p) {
		t.Fatal("should return true")
	}
}
