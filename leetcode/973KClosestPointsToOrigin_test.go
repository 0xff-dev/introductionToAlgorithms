package leetcode

import "testing"

func TestKClosest(t *testing.T) {
	points := [][]int{
		{1, 3},
		{-2, 2},
	}

	r := kClosest(points, 1)
	t.Logf("%v ", r)

	points = [][]int{
		{3, 3},
		{5, -1},
		{-2, 4},
	}

	r = kClosest(points, 3)
	t.Logf("%v ", r)
}
