package leetcode

import "testing"

func TestMostPoints(t *testing.T) {
	q := [][]int{
		{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5},
	}
	if r := mostPoints(q); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}

	q = [][]int{
		{12, 46}, {78, 19}, {63, 15}, {79, 62}, {13, 10},
	}
	if r := mostPoints(q); r != 79 {
		t.Fatalf("expect 79 get %d", r)
	}
}
