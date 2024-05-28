package leetcode

import "testing"

func TestEqualSubstring(t *testing.T) {
	ss, tt := "abcd", "bcdf"
	exp := 3
	if r := equalSubstring(ss, tt, 3); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	ss, tt = "abcd", "cdef"
	exp = 1
	if r := equalSubstring(ss, tt, 3); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	ss, tt = "abcd", "acde"
	exp = 1
	if r := equalSubstring(ss, tt, 0); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
