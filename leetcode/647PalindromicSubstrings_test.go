package leetcode

import "testing"

func TestCountSubstring(t *testing.T) {
	s := "abc"
	if r := countSubstrings(s); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	if r := countSubstringsDp(s); r != 3 {
		t.Fatalf("expect 3g get %d", r)
	}

	s = "aaa"
	if r := countSubstrings(s); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}
	if r := countSubstringsDp(s); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}

	s = "ababa"
	if r := countSubstringsDp(s); r != 9 {
		t.Fatalf("expect 9get %d", r)
	}

	s = "aaaaaa"
	if r := countSubstringsDp(s); r != 21 {
		t.Fatalf("expect 21 get %d", r)
	}
}
