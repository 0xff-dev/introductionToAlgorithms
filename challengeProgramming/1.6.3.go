/*
	一堆纸片有数字，从中选取4个，是否存在和为某个值
	当n较小的时候可以采取4层for循环，较大的时候，可以采取三层+二分
*/
package challengeProgramming

import "sort"

func binarySearch(aim int, nums []int) bool {
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] == aim {
			return true
		} else if nums[mid] < aim {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}
func FindSum(aim int, nums []int) bool {
	sort.Ints(nums)
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			for u := 0; u < length; u++ {
				tmp := nums[i] + nums[j] + nums[u]
				if binarySearch(aim-tmp, nums) {
					return true
				}
			}
		}
	}
	return false
}
