package leetcode

import "testing"

func TestSortColors(t *testing.T) {
	input := []int{2, 0, 2, 1, 1, 0}
	sortColors(input)
	t.Logf("%v", input)

	input = []int{0}
	sortColors(input)
	t.Logf("%v", input)

	input = []int{1, 0, 2}
	sortColors(input)
	t.Logf("%v", input)
}
