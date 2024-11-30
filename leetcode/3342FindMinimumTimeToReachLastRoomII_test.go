package leetcode

import "testing"

func TestMinTimeToReach3342(t *testing.T) {
	moveTime, exp := [][]int{{0, 4}, {4, 4}}, 7
	if r := minTimeToReach3342(moveTime); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	moveTime, exp = [][]int{{0, 0, 0, 0}, {0, 0, 0, 0}}, 6
	if r := minTimeToReach3342(moveTime); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	moveTime, exp = [][]int{{0, 1}, {1, 2}}, 4
	if r := minTimeToReach3342(moveTime); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
