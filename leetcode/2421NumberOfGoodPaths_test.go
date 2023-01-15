package leetcode

import "testing"

func TestNumberOfGoodPaths(t *testing.T) {
	vals := []int{1, 3, 2, 1, 3}
	edges := [][]int{
		{0, 1}, {0, 2}, {2, 3}, {2, 4},
	}
	if r := numberOfGoodPaths(vals, edges); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}

	vals = []int{0, 0, 0}
	edges = [][]int{
		{0, 1}, {0, 2},
	}
	if r := numberOfGoodPaths(vals, edges); r != 6 {
		t.Fatalf("expect 3 get %d", r)
	}

	vals = []int{1}
	edges = [][]int{}
	if r := numberOfGoodPaths(vals, edges); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	vals = []int{1, 1, 2, 2, 3}
	edges = [][]int{
		{0, 1}, {1, 2}, {2, 3}, {2, 4},
	}

	if r := numberOfGoodPaths(vals, edges); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}
}
