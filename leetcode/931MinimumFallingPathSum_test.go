package leetcode

import "testing"

func TestMinFallingPathSum1(t *testing.T) {
	matrix := [][]int{
		{2, 1, 3},
		{6, 5, 4},
		{7, 8, 9},
	}

	if r := minFallingPathSum1(matrix); r != 13 {
		t.Fatalf("expect 13 get %d", r)
	}

	matrix = [][]int{
		{-19, 57},
		{-40, -5},
	}

	if r := minFallingPathSum1(matrix); r != -59 {
		t.Fatalf("expect -59 get %d", r)
	}

	matrix = [][]int{
		{-48},
	}

	if r := minFallingPathSum1(matrix); r != -48 {
		t.Fatalf("expect -48 get %d", r)
	}
}
