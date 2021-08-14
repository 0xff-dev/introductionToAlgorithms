package leetcode

import "testing"

func TestMatrixBlockSum(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	r := matrixBlockSum(matrix, 1)
	t.Logf("%v", r)
	matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	r = matrixBlockSum(matrix, 2)
	t.Logf("%v", r)
}
