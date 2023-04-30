package leetcode

import "testing"

func TestMaxNumEdgesToRemove(t *testing.T) {
	n := 2
	edges := [][]int{
		{1, 1, 2}, {2, 1, 2}, {3, 1, 2},
	}
	if r := maxNumEdgesToRemove(n, edges); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
