package challengeprogrammingdatastructure

import "testing"

func TestRangeSearch(t *testing.T) {
	n := 6
	ps := [][]int{
		{2, 1}, {2, 2}, {4, 2}, {6, 2}, {3, 3}, {5, 4},
	}
	queries := [][]int{
		{2, 4, 0, 4},
		{4, 10, 2, 5},
	}
	RangeSearch(n, ps, queries)
}
