package leetcode

import "testing"

func TestLargestNumber(t *testing.T) {
	input := []int{10, 2}
	t.Log(largestNumber(input))

	input = []int{3, 30, 34, 5, 9}
	t.Log(largestNumber(input))

	input = []int{1}
	t.Log(largestNumber(input))

	input = []int{0, 0}
	t.Log(largestNumber(input))

	input = []int{1, 0, 28, 3, 0}
	t.Log(largestNumber(input))
}
