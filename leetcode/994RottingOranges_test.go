package leetcode

import "testing"

func TestOrangesRotting(t *testing.T) {
	grid := [][]int{
		{2, 1, 1}, {1, 1, 0}, {0, 1, 1},
	}
	exp := 4
	if r := orangesRotting(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	grid = [][]int{
		{2, 1, 1}, {0, 1, 1}, {1, 0, 1},
	}
	exp = -1
	if r := orangesRotting(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
