package leetcode

import "testing"

func TestMaxWidthOfVerticalArea(t *testing.T) {
	input := [][]int{{8, 7}, {9, 9}, {7, 4}, {9, 7}}
	r := maxWidthOfVerticalArea(input)
	t.Log(r)

	input = [][]int{{3, 1}, {9, 0}, {1, 0}, {1, 4}, {5, 3}, {8, 8}}
	r = maxWidthOfVerticalArea(input)
	t.Log(r)
}
