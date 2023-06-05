package leetcode

import "testing"

func TestCheckStraightLine(t *testing.T) {
	coordinates := [][]int{
		{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7},
	}
	if !checkStraightLine(coordinates) {
		t.Fatalf("expect true get false")
	}
	coordinates = [][]int{
		{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7},
	}
	if checkStraightLine(coordinates) {
		t.Fatalf("expect false get true")
	}

	coordinates = [][]int{
		{1, 1}, {2, 2}, {2, 0},
	}
	if checkStraightLine(coordinates) {
		t.Fatalf("expect false get true")
	}
}
