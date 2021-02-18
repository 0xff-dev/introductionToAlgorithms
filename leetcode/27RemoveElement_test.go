package leetcode

import "testing"

func TestRemoveElement(t *testing.T) {
	input := []int{3, 2, 2, 3}
	r := removeElement(input, 3)
	t.Logf("length: %d elements: %v", r, input[:r])

	input = []int{0, 1, 2, 2, 3, 0, 4, 2}
	r = removeElement(input, 2)
	t.Logf("length: %d elements: %v", r, input[:r])
}
