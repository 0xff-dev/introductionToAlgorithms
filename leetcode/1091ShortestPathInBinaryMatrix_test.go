package leetcode

import "testing"

func TestShortestPathBinaryMatrix(t *testing.T) {
	grid := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}
	if r := shortestPathBinaryMatrix(grid); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	grid = [][]int{
		{0, 1}, {1, 0},
	}
	if r := shortestPathBinaryMatrix(grid); r != 2 {
		t.Fatalf("expect 4 get %d", r)
	}
	grid = [][]int{
		{0, 0, 1, 0, 0, 0, 0}, {0, 1, 0, 0, 0, 0, 1}, {0, 0, 1, 0, 1, 0, 0}, {0, 0, 0, 1, 1, 1, 0}, {1, 0, 0, 1, 1, 0, 0}, {1, 1, 1, 1, 1, 0, 1}, {0, 0, 1, 0, 0, 0, 0},
	}
	if r := shortestPathBinaryMatrix(grid); r != 10 {
		t.Fatalf("expect 10 get %d", r)
	}
}
