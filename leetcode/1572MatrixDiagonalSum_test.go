package leetcode

import "testing"

func TestDiagonalSum(t *testing.T) {
	mat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	if r := diagonalSum(mat); r != 25 {
		t.Fatalf("expect 25 get %d", r)
	}
	mat = [][]int{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	}
	if r := diagonalSum(mat); r != 8 {
		t.Fatalf("expect 8 get %d", r)
	}
}
