package leetcode

import "testing"

func TestCountPaths1976(t *testing.T) {
	n, roads, exp := 7, [][]int{{0, 6, 7}, {0, 1, 2}, {1, 2, 3}, {1, 3, 3}, {6, 3, 3}, {3, 5, 1}, {6, 5, 1}, {2, 5, 1}, {0, 4, 5}, {4, 6, 2}}, 4
	if r := countPaths1976(n, roads); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
