package leetcode

import "testing"

func TestMinTimeToReach(t *testing.T) {
	moveTime, exp := [][]int{{0, 4}, {4, 4}}, 6
	if r := minTimeToReach(moveTime); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	moveTime, exp = [][]int{{0, 0, 0}, {0, 0, 0}}, 3
	if r := minTimeToReach(moveTime); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	moveTime, exp = [][]int{{0, 1}, {1, 2}}, 3
	if r := minTimeToReach(moveTime); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
