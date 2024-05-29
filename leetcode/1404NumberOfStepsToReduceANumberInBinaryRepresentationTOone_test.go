package leetcode

import "testing"

func TestNumSteps(t *testing.T) {
	s := "1101"
	exp := 6
	if r := numSteps(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	s = "10"
	exp = 1
	if r := numSteps(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
