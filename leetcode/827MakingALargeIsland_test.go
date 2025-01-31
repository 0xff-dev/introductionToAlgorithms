package leetcode

import "testing"

func TestLargestIsland(t *testing.T) {
	grid, exp := [][]int{{1, 0}, {0, 1}}, 3
	if r := largestIsland(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	grid, exp = [][]int{{1, 1}, {1, 0}}, 4
	if r := largestIsland(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	grid, exp = [][]int{{1, 1}, {1, 1}}, 4
	if r := largestIsland(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
