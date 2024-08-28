package leetcode

import "testing"

func TestCountSubIslands(t *testing.T) {
	g1, g2, exp := [][]int{{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}}, [][]int{{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 0}}, 3
	if r := countSubIslands(g1, g2); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	g1, g2, exp = [][]int{{1, 0, 1, 0, 1}, {1, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {1, 0, 1, 0, 1}}, [][]int{{0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {0, 1, 0, 1, 0}, {0, 1, 0, 1, 0}, {1, 0, 0, 0, 1}}, 2
	if r := countSubIslands(g1, g2); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
