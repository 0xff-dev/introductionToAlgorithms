package challengeprogrammingdatastructure

import "testing"

func TestConnectedComonents(t *testing.T) {
	n := 10
	graph := [][]int{
		{0, 1}, {0, 2}, {3, 4}, {5, 7}, {5, 6}, {6, 7}, {6, 8}, {7, 8}, {8, 9},
	}
	tc := [][]int{
		{0, 1}, {5, 9}, {1, 3},
	}
	ConnectedComponents(n, graph, tc)
}
