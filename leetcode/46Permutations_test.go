package leetcode

import "testing"

func TestPermute(t *testing.T) {
	input := []int{1, 2, 3}
	r := permute(input)
	t.Log(r)

	input = []int{1}
	r = permute(input)
	t.Log(r)

	input = []int{1, 0}
	r = permute(input)
	t.Log(r)

	input = []int{1, 1, 2}
	r = permute(input)
	t.Log(r)
}
