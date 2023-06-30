package leetcode

import "testing"

func TestLatestDayToCross(t *testing.T) {
	row, col := 2, 2
	cells := [][]int{
		{1, 1}, {2, 1}, {1, 2}, {2, 2},
	}
	if r := latestDayToCross(row, col, cells); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	row, col = 2, 2
	cells = [][]int{
		{1, 1}, {1, 2}, {2, 1}, {2, 2},
	}
	if r := latestDayToCross(row, col, cells); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	row, col = 3, 3
	cells = [][]int{
		{1, 2}, {2, 1}, {3, 3}, {2, 2}, {1, 1}, {1, 3}, {2, 3}, {3, 2}, {3, 1},
	}
	if r := latestDayToCross(row, col, cells); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	row, col = 3, 4
	cells = [][]int{
		{3, 1}, {2, 3}, {1, 3},
		{3, 2}, {2, 1}, {1, 4},
		{1, 1}, {2, 4}, {3, 4},
		{1, 2}, {2, 2}, {3, 3},
	}
	if r := latestDayToCross(row, col, cells); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}
}
