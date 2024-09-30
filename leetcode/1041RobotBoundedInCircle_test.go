package leetcode

import "testing"

func TestIsRobotBounded(t *testing.T) {
	s, exp := "GGLLGG", true
	if r := isRobotBounded(s); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
	s, exp = "GG", false
	if r := isRobotBounded(s); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
	s, exp = "GL", true
	if r := isRobotBounded(s); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
