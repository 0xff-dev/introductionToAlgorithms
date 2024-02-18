package leetcode

import "testing"

func TestMostBooked(t *testing.T) {
	n := 3
	meetings := [][]int{
		{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8},
	}
	if r := mostBooked(n, meetings); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
