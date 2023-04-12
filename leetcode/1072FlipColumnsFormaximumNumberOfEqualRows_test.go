package leetcode

import "testing"

func TestMaxEqualRowsAfterFlips(t *testing.T) {
	matrix := [][]int{
		{0, 1}, {1, 1},
	}
	if r := maxEqualRowsAfterFlips(matrix); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	matrix = [][]int{
		{0, 1}, {1, 0},
	}
	if r := maxEqualRowsAfterFlips(matrix); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	matrix = [][]int{
		{0, 0, 0}, {0, 0, 1}, {1, 1, 0},
	}
	if r := maxEqualRowsAfterFlips(matrix); r != 2 {
		t.Fatalf("expect 3 get %d", r)
	}
}
