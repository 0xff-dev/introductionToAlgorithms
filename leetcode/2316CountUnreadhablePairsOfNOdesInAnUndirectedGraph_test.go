package leetcode

import "testing"

func TestCountPairs(t *testing.T) {
	n := 3
	edges := [][]int{
		{0, 1}, {0, 2}, {1, 2},
	}
	if r := countPairs(n, edges); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	n = 7
	edges = [][]int{
		{0, 2}, {0, 5}, {2, 4}, {1, 6}, {5, 4},
	}
	if r := countPairs(n, edges); r != 14 {
		t.Fatalf("expect 14 get %d", r)
	}
}
