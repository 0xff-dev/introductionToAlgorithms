package leetcode

import "testing"

func TestTwoSum(t *testing.T) {
	input, target := []int{2, 7, 11, 15}, 9
	r := twoSum(input, target)
	t.Log(r)

	input, target = []int{3, 2, 4}, 6
	r = twoSum(input, target)
	t.Log(r)

	input, target = []int{2}, 4
	r = twoSum(input, target)
	t.Log(r)

	input, target = []int{2, 2}, 4
	r = twoSum(input, target)
	t.Log(r)

	input, target = []int{3, 3}, 6
	r = twoSum(input, target)
	t.Log(r)
}
