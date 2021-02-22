package leetcode

import "testing"

func TestCombinationSum(t *testing.T) {
	input, target := []int{2, 3, 6, 7}, 7
	r := combinationSum(input, target)
	t.Log(r)

	input, target = []int{2, 3, 5}, 8
	r = combinationSum(input, target)
	t.Log(r)

	input, target = []int{2, 7, 6, 3, 5, 1}, 9
	t.Log(combinationSum(input, target))
}
