package leetcode

import "testing"

func TestColoredCells(t *testing.T) {
	n, exp := 1, int64(1)
	if r := coloredCells(n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	n, exp = 2, 5
	if r := coloredCells(n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	n, exp = 3, 13
	if r := coloredCells(n); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
