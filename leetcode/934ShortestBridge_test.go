package leetcode

import "testing"

func TestShortestBridge(t *testing.T) {
	grid := [][]int{
		{0, 1}, {1, 0},
	}
	if r := shortestBridge(grid); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	grid = [][]int{
		{0, 1, 0}, {0, 0, 0}, {0, 0, 1},
	}
	if r := shortestBridge(grid); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	grid = [][]int{
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
	}
	if r := shortestBridge(grid); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	grid = [][]int{
		{0, 0, 0, 1, 1},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 1, 1},
		{0, 0, 1, 0, 1},
		{0, 0, 1, 1, 0},
	}
	if r := shortestBridge(grid); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
