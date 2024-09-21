package leetcode

import "testing"

func TestCountSquares(t *testing.T) {
	matrix, exp := [][]int{{0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}}, 15
	if r := countSquares(matrix); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	matrix, exp = [][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}, 7
	if r := countSquares(matrix); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
