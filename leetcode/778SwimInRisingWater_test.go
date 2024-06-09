package leetcode

import "testing"

func TestSwimInWater(t *testing.T) {
	grid := [][]int{
		{0, 2}, {1, 3},
	}
	exp := 3
	if r := swimInWater(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	grid = [][]int{
		{0, 1, 2, 3, 4},
		{24, 23, 22, 21, 5},
		{12, 13, 14, 15, 16},
		{11, 17, 18, 19, 20},
		{10, 9, 8, 7, 6},
	}
	exp = 16
	if r := swimInWater(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
