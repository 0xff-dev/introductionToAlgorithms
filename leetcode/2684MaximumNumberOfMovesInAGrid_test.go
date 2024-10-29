package leetcode

import "testing"

func TestMaxMoves(t *testing.T) {
	grid, exp := [][]int{{2, 3, 4, 5}, {5, 4, 9, 3}, {3, 4, 2, 11}, {10, 9, 13, 15}}, 3
	if r := maxMoves(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	grid, exp = [][]int{{3, 2, 4}, {2, 1, 9}, {1, 1, 7}}, 0
	if r := maxMoves(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
