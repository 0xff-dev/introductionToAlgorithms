package leetcode

import "testing"

func TestFindMinArrowShots(t *testing.T) {
	points := [][]int{
		{10, 16}, {2, 8}, {1, 6}, {7, 12},
	}
	if r := findMinArrowShots(points); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	points = [][]int{
		{1, 2}, {3, 4}, {5, 6}, {7, 8},
	}

	if r := findMinArrowShots(points); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	points = [][]int{
		{1, 2}, {2, 3}, {3, 4}, {4, 5},
	}

	if r := findMinArrowShots(points); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
