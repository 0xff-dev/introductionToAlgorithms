package leetcode

import "testing"

func TestClosestMeetingNode(t *testing.T) {
	e := []int{2, 2, 3, -1}
	n1, n2 := 0, 1
	if r := closestMeetingNode(e, n1, n2); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
