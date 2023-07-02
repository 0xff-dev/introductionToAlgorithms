package leetcode

import "testing"

func TestMaximumRequests(t *testing.T) {
	n := 5
	requets := [][]int{
		{0, 1}, {1, 0}, {0, 1}, {1, 2}, {2, 0}, {3, 4},
	}
	if r := maximumRequests(n, requets); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}

	n = 3
	requets = [][]int{
		{0, 0}, {1, 2}, {2, 1},
	}
	if r := maximumRequests(n, requets); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	n = 4
	requets = [][]int{
		{0, 3}, {3, 1}, {1, 2}, {2, 0},
	}
	if r := maximumRequests(n, requets); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
}
