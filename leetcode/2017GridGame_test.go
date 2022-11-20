package leetcode

import "testing"

func TestGridGame(t *testing.T) {
	grid := [][]int{
		{2, 5, 4},
		{1, 5, 1},
	}
	if r := gridGame(grid); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	grid = [][]int{
		{3, 3, 1},
		{8, 5, 2},
	}
	if r := gridGame(grid); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	grid = [][]int{
		{1, 3, 1, 15},
		{1, 3, 3, 1},
	}

	if r := gridGame(grid); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}

	grid = [][]int{
		{2, 4, 6, 8, 10, 12},
		{11, 9, 7, 5, 3, 1},
	}

	if r := gridGame(grid); r != 27 {
		t.Fatalf("expect 27 get %d", r)
	}
}
