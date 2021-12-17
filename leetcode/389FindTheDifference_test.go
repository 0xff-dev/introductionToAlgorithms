package leetcode

import "testing"

func TestFindTheDifference(t *testing.T) {
	s, t1 := "abcd", "abcde"
	if r := findTheDifference(s, t1); r != 'e' {
		t.Fatalf("expect 'e' get %c", r)
	}

	s, t1 = "", "y"
	if r := findTheDifference(s, t1); r != 'y' {
		t.Fatalf("expect 'y' get %c", r)
	}

	s, t1 = "a", "aa"
	if r := findTheDifference(s, t1); r != 'a' {
		t.Fatalf("expect 'a' get %c", r)
	}

	s, t1 = "ae", "aea"
	if r := findTheDifference(s, t1); r != 'a' {
		t.Fatalf("expect 'a' get %c", r)
	}
}
