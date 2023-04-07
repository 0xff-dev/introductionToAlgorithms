package leetcode

import "testing"

func TestNumEnclaves(t *testing.T) {
	grid := [][]int{
		{0, 0, 0, 0},
		{1, 0, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	}
	if r := numEnclaves(grid); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	grid = [][]int{
		{0, 1, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 0},
	}
	if r := numEnclaves(grid); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}
