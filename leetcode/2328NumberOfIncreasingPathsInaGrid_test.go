package leetcode

import "testing"

func TestCountPaths(t *testing.T) {
	grid := [][]int{
		{1, 1}, {3, 4},
	}
	if r := countPaths(grid); r != 8 {
		t.Fatalf("expect 8 get %d", r)
	}
	grid = [][]int{
		{1, 2},
	}
	if r := countPaths(grid); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}
