package leetcode

import "testing"

func TestCountCompleteComponents(t *testing.T) {
	n := 6
	edges := [][]int{
		{0, 1}, {0, 2}, {1, 2}, {3, 4},
	}
	if r := countCompleteComponents(n, edges); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	n = 6
	edges = [][]int{
		{0, 1}, {0, 2}, {1, 2}, {3, 4}, {3, 5},
	}
	if r := countCompleteComponents(n, edges); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
