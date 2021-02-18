package leetcode

import "testing"

func TestNextPermutation(t *testing.T) {
	input := []int{1, 2, 3}
	nextPermutation(input)
	t.Logf("%v", input)

	input = []int{3, 2, 1}
	nextPermutation(input)
	t.Logf("%v", input)

	input = []int{1, 4, 3, 8, 2}
	nextPermutation(input)
	t.Logf("%v", input)

	input = []int{1}
	nextPermutation(input)
	t.Logf("%v", input)

	input = []int{4, 8, 6, 6, 4}
	nextPermutation(input)
	t.Logf("%v", input)
}
