package leetcode

import "testing"

func TestMinimumLength(t *testing.T) {
	s := "bbbbbbbbbbbbbbbbbbbbbbbbbbbabbbbbbbbbbbbbbbccbcbcbccbbabbb"
	exp := 1
	if r := minimumLength(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	s = "cabaabac"
	exp = 0
	if r := minimumLength(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
