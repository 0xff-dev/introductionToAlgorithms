package leetcode

import "testing"

func TestMinCost1368(t *testing.T) {
	grid, exp := [][]int{{1, 1, 1, 1}, {2, 2, 2, 2}, {1, 1, 1, 1}, {2, 2, 2, 2}}, 3
	if r := minCost1368(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	grid, exp = [][]int{{1, 1, 3}, {3, 2, 2}, {1, 1, 4}}, 0
	if r := minCost1368(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	grid, exp = [][]int{{1, 2}, {4, 3}}, 1
	if r := minCost1368(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
