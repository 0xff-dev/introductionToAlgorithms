package leetcode

import "testing"

func TestMinimumEffortPath(t *testing.T) {
	heights := [][]int{
		{1, 2, 2},
		{3, 8, 2},
		{5, 3, 5},
	}
	if r := minimumEffortPath(heights); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
