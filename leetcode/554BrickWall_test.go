package leetcode

import "testing"

func TestLeastBricks(t *testing.T) {
	grid, exp := [][]int{{1, 2, 2, 1}, {3, 1, 2}, {1, 3, 2}, {2, 4}, {3, 1, 2}, {1, 3, 1, 1}}, 2
	if r := leastBricks(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	grid, exp = [][]int{{1}, {1}, {1}}, 3
	if r := leastBricks(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	grid, exp = [][]int{{100000000}, {100000000}, {100000000}}, 3
	if r := leastBricks(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
