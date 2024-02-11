package leetcode

import "testing"

func TestCherryPickup(t *testing.T) {
	grid := [][]int{
		{3, 1, 1}, {2, 5, 1}, {1, 5, 5}, {2, 1, 1},
	}
	exp := 24
	if r := cherryPickup(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	grid = [][]int{
		{1, 0, 0, 0, 0, 0, 1},
		{2, 0, 0, 0, 0, 3, 0},
		{2, 0, 9, 0, 0, 0, 0},
		{0, 3, 0, 5, 4, 0, 0},
		{1, 0, 2, 3, 0, 0, 6},
	}
	exp = 28
	if r := cherryPickup(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
