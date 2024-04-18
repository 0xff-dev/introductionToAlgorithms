package leetcode

import "testing"

func TestMinSteps650(t *testing.T) {
	n := 3
	exp := 3
	if r := minSteps650(n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	n = 1
	exp = 0
	if r := minSteps650(n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
