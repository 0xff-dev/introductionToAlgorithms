package leetcode

import "testing"

func TestEraseOverlapIntervals(t *testing.T) {
	intervals := [][]int{
		{1, 2}, {2, 3}, {3, 4}, {1, 3},
	}
	if r := eraseOverlapIntervals(intervals); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	intervals = [][]int{
		{1, 2}, {1, 2}, {1, 2},
	}
	if r := eraseOverlapIntervals(intervals); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	intervals = [][]int{
		{1, 2}, {2, 3},
	}
	if r := eraseOverlapIntervals(intervals); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}
