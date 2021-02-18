package leetcode

import (
	"testing"
)

func TestRemoveDuplicate(t *testing.T) {
	input := []int{1, 1, 2}
	r := removeDuplicates(input)
	t.Logf("length: %d v: %v", r, input[:r])

	input = []int{1, 1, 1, 1, 1}
	r = removeDuplicates(input)
	t.Logf("length: %d v: %v", r, input[:r])

	input = []int{1}
	r = removeDuplicates(input)
	t.Logf("length: %d v: %v", r, input[:r])

	input = []int{1, 2}
	r = removeDuplicates(input)
	t.Logf("length: %d v: %v", r, input[:r])

	input = []int{1, 2, 2, 2, 2, 2}
	r = removeDuplicates(input)
	t.Logf("length: %d v: %v", r, input[:r])

	input = []int{1, 1, 1, 1, 1, 2, 2, 2, 3, 4, 4, 4, 5}
	r = removeDuplicates(input)
	t.Logf("length: %d v: %v", r, input[:r])
}
