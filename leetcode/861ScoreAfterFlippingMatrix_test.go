package leetcode

import "testing"

// [[0,0,1,1],[1,0,1,0],[1,1,0,0]] 39
// 0 1
func TestMatrixScore(t *testing.T) {
	matrix := [][]int{
		{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0},
	}
	exp := 39
	if r := matrixScore(matrix); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	matrix = [][]int{{0}}
	exp = 1
	if r := matrixScore(matrix); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
