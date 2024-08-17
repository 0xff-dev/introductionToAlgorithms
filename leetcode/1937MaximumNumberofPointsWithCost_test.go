package leetcode

import "testing"

func TestMaxPoints1937(t *testing.T) {
	points, exp := [][]int{{1, 2, 3}, {1, 5, 1}, {3, 1, 1}}, int64(9)
	if r := maxPoints1937(points); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	points, exp = [][]int{{1, 5}, {2, 3}, {4, 2}}, 11
	if r := maxPoints1937(points); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
