package leetcode

import "testing"

func TestMaxUniqueSplit(t *testing.T) {
	s, exp := "ababccc", 5
	if r := maxUniqueSplit(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	s, exp = "aba", 2
	if r := maxUniqueSplit(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	s, exp = "aa", 1
	if r := maxUniqueSplit(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
