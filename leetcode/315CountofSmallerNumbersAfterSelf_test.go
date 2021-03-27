package leetcode

import "testing"

func TestCountSmaller(t *testing.T) {
	input := []int{5, 2, 6, 1}
	r := countSmaller(input)
	t.Logf("%v", r)

	input = []int{-1}
	r = countSmaller(input)
	t.Logf("%v", r)

	input = []int{-1, -1}
	r = countSmaller(input)
	t.Logf("%v", r)

	input = []int{5, 3, 5, 2, 1}
	r = countSmaller(input)
	t.Logf("%v", r)

	input = []int{5, 4, 3, 2, 1}
	r = countSmaller(input)
	t.Logf("%v", r)

	input = []int{1, 2, 3, 4, 5}
	r = countSmaller(input)
	t.Logf("%v", r)

	input = []int{8, 1, 3, 0, 10, 2, 8}
	r = countSmaller(input)
	t.Logf("%v", r)
}
