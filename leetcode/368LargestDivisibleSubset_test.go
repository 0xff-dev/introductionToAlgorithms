package leetcode

import "testing"

func TestLargestDivisibleSubset(t *testing.T) {
	input := []int{1, 2, 3}
	r := largestDivisibleSubset(input)
	t.Logf("%v", r)

	input = []int{1, 2, 4, 8}
	r = largestDivisibleSubset(input)
	t.Logf("%v", r)

	input = []int{9, 6, 3, 1}
	r = largestDivisibleSubset(input)
	t.Logf("%v", r)

	input = []int{1}
	r = largestDivisibleSubset(input)
	t.Logf("%v", r)

	input = []int{5, 9, 18, 54, 108, 540, 90, 180, 360, 720}
	r = largestDivisibleSubset(input)
	t.Logf("%v", r)
}
