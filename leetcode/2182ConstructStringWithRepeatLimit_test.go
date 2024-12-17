package leetcode

import "testing"

func TestRepeatLimitedString(t *testing.T) {
	s, r, exp := "cczazcc", 3, "zzcccac"
	if r := repeatLimitedString(s, r); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}

	s, r, exp = "aababab", 2, "bbabaa"
	if r := repeatLimitedString(s, r); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
}
