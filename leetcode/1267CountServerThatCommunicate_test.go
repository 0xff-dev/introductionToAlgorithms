package leetcode

import "testing"

func TestCountServers(t *testing.T) {
	grid := [][]int{
		{1, 0}, {0, 1},
	}
	if r := countServers(grid); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
	grid = [][]int{
		{1, 0}, {1, 1},
	}
	if r := countServers(grid); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	grid = [][]int{
		{1, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
	if r := countServers(grid); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
}
