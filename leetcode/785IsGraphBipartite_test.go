package leetcode

import "testing"

func TestIsBipartite(t *testing.T) {
	graph := [][]int{
		{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2},
	}
	if isBipartite(graph) {
		t.Fatalf("expect false get true")
	}
	graph = [][]int{
		{1, 3}, {0, 2}, {1, 3}, {0, 2},
	}
	if !isBipartite(graph) {
		t.Fatalf("expect true get false")
	}

	graph = [][]int{
		{3}, {2, 4}, {1}, {0, 4}, {1, 3},
	}
	if !isBipartite(graph) {
		t.Fatalf("expect true get false")
	}
}
