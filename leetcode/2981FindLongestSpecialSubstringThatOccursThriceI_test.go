package leetcode

import "testing"

func TestMaximumLength(t *testing.T) {
	s, exp := "aaaa", 2
	if r := maximumLength(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	s, exp = "abcdef", -1
	if r := maximumLength(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	s, exp = "abcaba", 1
	if r := maximumLength(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
