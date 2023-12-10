package leetcode

import "sort"

func triangleNumber(nums []int) int {
	length := len(nums)
	if length < 3 {
		return 0
	}

	var bsearch func(int, int, int, int) int
	bsearch = func(left, right, a, b int) int {
		for left < right {
			mid := left + (right-left)/2
			if nums[a]+nums[b] > nums[mid] {
				left = mid + 1
				continue
			}
			right = mid
		}
		return left
	}
	ans := 0
	sort.Ints(nums)
	for i := 0; i < length-2; i++ {
		for j := i + 1; j < length-1; j++ {
			index := bsearch(j+1, length, i, j)
			ans += index - j - 1
		}
	}
	return ans
}
