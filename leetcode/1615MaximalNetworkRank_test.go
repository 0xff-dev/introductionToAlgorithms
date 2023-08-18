package leetcode

import "testing"

func TestMaximalNetworkRank(t *testing.T) {
	n := 4
	roads := [][]int{
		{0, 1}, {0, 3}, {1, 2}, {1, 3},
	}
	if r := maximalNetworkRank(n, roads); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
	n = 5
	roads = [][]int{
		{0, 1}, {0, 3}, {1, 2}, {1, 3}, {2, 3}, {2, 4},
	}
	if r := maximalNetworkRank(n, roads); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}
	n = 8
	roads = [][]int{
		{0, 1}, {1, 2}, {2, 3}, {2, 4}, {5, 6}, {5, 7},
	}
	if r := maximalNetworkRank(n, roads); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}
}
