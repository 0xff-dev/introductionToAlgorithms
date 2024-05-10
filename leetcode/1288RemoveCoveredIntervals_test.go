package leetcode

import "testing"

func TestRemoveCoveredIntervals(t *testing.T) {
	intervals := [][]int{
		{1, 4}, {3, 6}, {2, 8},
	}
	exp := 2
	if r := removeCoveredIntervals(intervals); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	intervals = [][]int{{1, 4}, {2, 3}}
	exp = 1
	if r := removeCoveredIntervals(intervals); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
