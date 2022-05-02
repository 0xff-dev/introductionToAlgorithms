package leetcode

import "testing"

func TestSortArrayInParity(t *testing.T) {
	nums := []int{3, 1, 2, 4}
	r := sortArrayByParity(nums)
	t.Logf("%v", r)

	nums = []int{1, 3, 5, 7, 9, 2, 4, 8, 11}
	t.Logf("%v", sortArrayByParity(nums))
}
