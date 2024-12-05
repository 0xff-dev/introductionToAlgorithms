package leetcode

import "testing"

func TestCanChange(t *testing.T) {
	start, target, exp := "_L__R__R_", "L______RR", true
	if r := canChange(start, target); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
	start, target, exp = "R_L_", "__LR", false
	if r := canChange(start, target); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
	start, target, exp = "_R", "R_", false
	if r := canChange(start, target); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
