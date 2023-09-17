package leetcode

import "testing"

func TestShortestPathLength(t *testing.T) {
	graph := [][]int{
		{1, 2, 3}, {0}, {0}, {0},
	}
	if r := shortestPathLength(graph); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
}
