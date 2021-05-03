package leetcode

import "testing"

func TestMinFallingPathSum(t *testing.T) {
	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	if r := minFallingPathSum(grid); r != 13 {
		t.Fatalf("expect 13 get %d", r)
	}

}
