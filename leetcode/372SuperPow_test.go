package leetcode

import "testing"

func TestSuperPow(t *testing.T) {
	a, b := 2, []int{3}
	exp := 8
	if r := superPow(a, b); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	a, b = 2, []int{1, 0}
	exp = 1024
	if r := superPow(a, b); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	a, b = 15, []int{1, 2, 3}
	exp = 386
	if r := superPow(a, b); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
