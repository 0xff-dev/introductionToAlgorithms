package leetcode

import "testing"

func TestIsLandPerimeter(t *testing.T) {
	grid := [][]int{
		{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0},
	}
	exp := 16
	if r := islandPerimeter(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	grid = [][]int{
		{1},
	}
	exp = 4
	if r := islandPerimeter(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	grid = [][]int{
		{1, 0},
	}
	if r := islandPerimeter(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
