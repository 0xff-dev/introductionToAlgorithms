package leetcode

import "testing"

func TestMinimumSteps(t *testing.T) {
	s, exp := "101", int64(1)
	if r := minimumSteps(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	s, exp = "100", 2
	if r := minimumSteps(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	s, exp = "0111", 0
	if r := minimumSteps(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
