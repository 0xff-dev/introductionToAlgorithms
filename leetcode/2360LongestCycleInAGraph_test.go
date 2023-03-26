package leetcode

import "testing"

func TestLongestCycle(t *testing.T) {
	edges := []int{3, 3, 4, 2, 3}
	if r := longestCycle(edges); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	edges = []int{-1, 4, -1, 2, 0, 4}
	if r := longestCycle(edges); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}
}
