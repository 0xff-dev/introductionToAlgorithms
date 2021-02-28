package leetcode

import "testing"

func TestNQueens(t *testing.T) {
	n := 4
	r := solveNQueens(n)
	for _, line := range r {
		t.Logf("%v\n", line)
	}

	n = 1
	r = solveNQueens(n)
	for _, line := range r {
		t.Logf("%v", line)
	}

	n = 8
	r = solveNQueens(n)
	for _, line := range r {
		t.Logf("%v", line)
	}
}
