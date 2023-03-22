package leetcode

import "testing"

func TestMinScore(t *testing.T) {
	n := 4
	roads := [][]int{
		{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7},
	}
	if r := minScore(n, roads); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}

	n = 4
	roads = [][]int{
		{1, 2, 2}, {1, 3, 4}, {3, 4, 7},
	}
	if r := minScore(n, roads); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
