package leetcode

import "testing"

func TestLargestPathValue(t *testing.T) {
	colors := "abaca"
	edges := [][]int{
		{0, 1}, {0, 2}, {2, 3}, {3, 4},
	}
	if r := largestPathValue(colors, edges); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	colors = "a"
	edges = [][]int{
		{0, 0},
	}
	if r := largestPathValue(colors, edges); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}

	colors = "vwnnvwvqvq"
	edges = [][]int{
		{0, 1}, {0, 2}, {2, 3},
		{3, 4}, {4, 5}, {5, 6},
		{5, 7}, {7, 8}, {6, 9},
	}
	if r := largestPathValue(colors, edges); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}
