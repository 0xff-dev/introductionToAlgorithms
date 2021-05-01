package leetcode

import "testing"

func TestNumIslands(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	if r := numIslands(grid); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	grid = [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	if r := numIslands(grid); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	grid = [][]byte{
		{'0'},
	}
	if r := numIslands(grid); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
	grid = [][]byte{
		{'1'},
	}
	if r := numIslands(grid); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	grid = [][]byte{
		{'1', '1', '0', '0'},
		{'0', '1', '0', '1'},
	}
	if r := numIslands(grid); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
