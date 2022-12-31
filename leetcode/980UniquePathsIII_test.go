package leetcode

import "testing"

func TestUniquePathsII(t *testing.T) {
	grid := [][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 2, -1},
	}

	if r := uniquePathsIII(grid); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	grid = [][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 2},
	}
	if r := uniquePathsIII(grid); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	grid = [][]int{
		{0, 1},
		{2, 0},
	}
	if r := uniquePathsIII(grid); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}
