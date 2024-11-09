package leetcode

import "testing"

func TestMinEnd(t *testing.T) {
	n, x, exp := 3, 4, int64(6)
	if r := minEnd(n, x); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	n, x, exp = 2, 7, 15
	if r := minEnd(n, x); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

}
