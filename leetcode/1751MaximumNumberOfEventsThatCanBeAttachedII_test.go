package leetcode

import "testing"

func TestMaxValue1751(t *testing.T) {
	events := [][]int{
		{1, 2, 4}, {3, 4, 3}, {2, 3, 1},
	}
	k := 2
	if r := maxValue1751(events, k); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}
	events = [][]int{
		{1, 2, 4}, {3, 4, 3}, {2, 3, 10},
	}
	k = 2
	if r := maxValue1751(events, k); r != 10 {
		t.Fatalf("expect 10 get %d", r)
	}

	events = [][]int{
		{1, 1, 1}, {2, 2, 2}, {3, 3, 3}, {4, 4, 4},
	}
	k = 3
	if r := maxValue1751(events, k); r != 9 {
		t.Fatalf("expect 9 get %d", r)
	}
}
