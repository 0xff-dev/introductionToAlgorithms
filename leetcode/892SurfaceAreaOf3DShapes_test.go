package leetcode

import "testing"

func TestSurfaceArea(t *testing.T) {
	grid := [][]int{
		{1, 2},
		{3, 4},
	}

	if r := surfaceArea(grid); r != 34 {
		t.Fatalf("expect 34 get %d", r)
	}

	grid = [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	if r := surfaceArea(grid); r != 32 {
		t.Fatalf("expect 32 get %d", r)
	}

	grid = [][]int{
		{2, 2, 2},
		{2, 1, 2},
		{2, 2, 2},
	}
	if r := surfaceArea(grid); r != 46 {
		t.Fatalf("expect 46 get %d", r)
	}

	grid = [][]int{
		{1},
	}

	if r := surfaceArea(grid); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}

	grid = [][]int{
		{2},
	}
	if r := surfaceArea(grid); r != 10 {
		t.Fatalf("expect 10 get %d", r)
	}
}
