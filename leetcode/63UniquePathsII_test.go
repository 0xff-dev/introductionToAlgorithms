package leetcode

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	grid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	if r := uniquePathsWithObstacles(grid); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	grid = [][]int{
		{0, 1},
		{0, 0},
	}
	if r := uniquePathsWithObstacles(grid); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	grid = [][]int{
		{0, 0},
	}

	if r := uniquePathsWithObstacles(grid); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	grid = [][]int{
		{0},
		{0},
	}

	if r := uniquePathsWithObstacles(grid); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	grid = [][]int{
		{0, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 0},
	}
	if r := uniquePathsWithObstacles(grid); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}
}
