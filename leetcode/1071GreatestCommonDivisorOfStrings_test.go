package leetcode

import "testing"

func TestGcdOfStrings(t *testing.T) {
	s1, s2 := "ABCABC", "ABC"
	if r := gcdOfStrings(s1, s2); r != "ABC" {
		t.Fatalf("expect 'ABC' get %s", r)
	}

	s1, s2 = "ABABAB", "ABAB"
	if r := gcdOfStrings(s1, s2); r != "AB" {
		t.Fatalf("expect 'AB' get %s", r)
	}

	s1, s2 = "LEET", "CODE"
	if r := gcdOfStrings(s1, s2); r != "" {
		t.Fatalf("expect '' get %s", r)
	}

}
