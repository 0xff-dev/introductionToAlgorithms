package leetcode

import "testing"

func TestNumMagicSquareInside(t *testing.T) {
	grid, exp := [][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}}, 1
	if r := numMagicSquaresInside(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	grid, exp = [][]int{{8}}, 0
	if r := numMagicSquaresInside(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
