package leetcode

import "testing"

func TestBeautifulSubstrings(t *testing.T) {
	s, k, exp := "baeyh", 2, 2
	if r := beautifulSubstrings(s, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	s, k, exp = "abba", 1, 3
	if r := beautifulSubstrings(s, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	s, k, exp = "bcdf", 1, 0
	if r := beautifulSubstrings(s, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
