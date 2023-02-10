package leetcode

import "testing"

func TestMaxDistance1162(t *testing.T) {
	grid := [][]int{
		{1, 0, 1}, {0, 0, 0}, {1, 0, 1},
	}
	if r := maxDistance1162(grid); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	grid = [][]int{
		{1, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	if r := maxDistance1162(grid); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
}
