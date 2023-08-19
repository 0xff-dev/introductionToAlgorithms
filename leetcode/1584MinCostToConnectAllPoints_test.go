package leetcode

import "testing"

func TestMinCostConnectPoints(t *testing.T) {
	points := [][]int{
		{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0},
	}
	if r := minCostConnectPoints(points); r != 20 {
		t.Fatalf("expect 20 get %d", r)
	}
	points = [][]int{
		{3, 12}, {-2, 5}, {-4, 1},
	}
	if r := minCostConnectPoints(points); r != 18 {
		t.Fatalf("expect 18 get %d", r)
	}
}
