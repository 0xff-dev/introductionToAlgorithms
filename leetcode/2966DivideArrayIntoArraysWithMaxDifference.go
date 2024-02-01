package leetcode

import "sort"

func divideArray2966(nums []int, k int) [][]int {
	length := len(nums)
	if length%3 != 0 {
		return nil
	}
	sort.Ints(nums)
	ans := make([][]int, 0)
	c := 0
	for i := 0; i < len(nums); i, c = i+1, (c+1)%3 {
		if c == 0 {
			ans = append(ans, []int{nums[i]})
			continue
		}
		ll := len(ans) - 1
		if nums[i]-ans[ll][0] > k {
			return nil
		}
		ans[ll] = append(ans[ll], nums[i])
	}
	return ans
}
