package leetcode

import "testing"

func TestMaxTwoEvents(t *testing.T) {
	events, exp := [][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}}, 4
	if r := maxTwoEvents(events); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	events, exp = [][]int{{1, 3, 2}, {4, 5, 2}, {1, 5, 5}}, 5
	if r := maxTwoEvents(events); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	events, exp = [][]int{{1, 5, 3}, {1, 5, 1}, {6, 6, 5}}, 8
	if r := maxTwoEvents(events); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
