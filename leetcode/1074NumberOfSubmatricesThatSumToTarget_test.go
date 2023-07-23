package leetcode

import "testing"

func TestNumSubmatrixSumTarget(t *testing.T) {
	matrix := [][]int{
		{0, 1, 0}, {1, 1, 1}, {0, 1, 0},
	}
	if r := numSubmatrixSumTarget(matrix, 0); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
}
