package leetcode

import "testing"

func TestSearchRange(t *testing.T) {
	input := []int{5, 7, 7, 8, 8, 10}
	r := searchRange(input, 8)
	t.Logf("%v", r)

	r = searchRange(input, 6)
	t.Logf("%v", r)

	input = []int{}
	r = searchRange(input, 0)
	t.Logf("%v", r)

	input = []int{1}
	r = searchRange(input, 1)
	t.Logf("%v", r)

	input = []int{1, 1, 1, 1}
	r = searchRange(input, 1)
	t.Logf("%v", r)

	input = []int{1, 1, 2, 2, 2, 3}
	r = searchRange(input, 2)
	t.Logf("%v", r)
}
