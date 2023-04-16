package challengeprogrammingdatastructure

import "testing"

func TestMimimumSpanningTreePaths(t *testing.T) {
	n := 6
	edges := [][]int{
		{0, 1, 1},
		{0, 2, 3},
		{1, 2, 1},
		{1, 3, 7},
		{2, 4, 1},
		{1, 4, 3},
		{3, 4, 1},
		{3, 5, 1},
		{4, 5, 6},
	}
	if r := MinimumSpanningTreeEdges(n, edges); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}
}
