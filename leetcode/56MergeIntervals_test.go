package leetcode

import "testing"

func TestMerge(t *testing.T) {
	input := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	r := merge(input)
	t.Logf("%v", r)

	input = [][]int{{1, 4}, {4, 5}}
	r = merge(input)
	t.Logf("%v", r)

	input = [][]int{{1, 3}}
	r = merge(input)
	t.Logf("%v", r)

	input = [][]int{{2, 4}, {1, 8}, {1, 5}}
	r = merge(input)
	t.Logf("%v", r)

	input = [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}
	r = merge(input)
	t.Logf("%v", r)
}
