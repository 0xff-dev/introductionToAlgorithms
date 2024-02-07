package leetcode

import "testing"

func TestNthSuperUglyNumber(t *testing.T) {
	n := 12
	p := []int{2, 7, 13, 19}
	exp := 32
	if r := nthSuperUglyNumber(n, p); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	n = 1
	p = []int{2, 3, 5}
	exp = 1
	if r := nthSuperUglyNumber(n, p); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
