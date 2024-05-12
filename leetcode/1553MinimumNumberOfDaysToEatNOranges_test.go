package leetcode

import "testing"

func TestMinDays(t *testing.T) {
	n := 10
	exp := 4
	if r := minDays(n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	n = 6
	exp = 3
	if r := minDays(n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
