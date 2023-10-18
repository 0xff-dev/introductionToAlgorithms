package leetcode

import "testing"

func TestMinimumTime2050(t *testing.T) {
	n := 3
	relations := [][]int{
		{1, 3}, {2, 3},
	}
	time := []int{3, 2, 5}
	if r := minimumTime2050(n, relations, time); r != 8 {
		t.Fatalf("expect 8 get %d", r)
	}

	n = 5
	relations = [][]int{
		{1, 5}, {2, 5}, {3, 5}, {3, 4}, {4, 5},
	}
	time = []int{1, 2, 3, 4, 5}
	if r := minimumTime2050(n, relations, time); r != 12 {
		t.Fatalf("expect 12 get %d", r)
	}
}
