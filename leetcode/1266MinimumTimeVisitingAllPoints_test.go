package leetcode

import "testing"

func TestMinTimeToVisitAllPoints(t *testing.T) {
	points := [][]int{{1, 1}, {3, 4}, {-1, 0}}
	if r := minTimeToVisitAllPoints(points); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}
	points = [][]int{{3, 2}, {-2, 2}}
	if r := minTimeToVisitAllPoints(points); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}
}
