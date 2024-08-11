package leetcode

import "testing"

func TestMinDays1568(t *testing.T) {
	grid, exp := [][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}, 2
	if r := minDays1568(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	grid, exp = [][]int{{1, 1}}, 2
	if r := minDays1568(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
