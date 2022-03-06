package leetcode

import "testing"

func TestTwoSum167(t *testing.T) {
	numbers := []int{2, 7, 11, 15}
	r := twoSum167(numbers, 9)
	r1 := twoSum_1(numbers, 9)
	t.Logf("%v-%v", r, r1)

	numbers = []int{2, 3, 4}
	r = twoSum167(numbers, 6)
	r1 = twoSum_1(numbers, 6)
	t.Logf("%v-%v", r, r1)

	numbers = []int{-1, 0}
	r = twoSum167(numbers, -1)
	r1 = twoSum_1(numbers, -1)
	t.Logf("%v-%v", r, r1)

	numbers = []int{2, 3, 7, 15}
	r = twoSum167(numbers, 9)
	r1 = twoSum_1(numbers, 9)
	t.Logf("%v-%v", r, r1)
}
