package leetcode

// 长度为n
// [-3,-1, 1, 2, 3, 20] 缺失的是我们知道缺失的是4
// -  -  1  2  3  20 6-2=4 <= 4
// 0  1  2, 3, 4, 5 长度是6 如果
//
//	1, 2, 3, 4 期待正数部署应该是1，2， 3，4，
//
// [ 1, 2, 3, 4, 5, 6] 正常应该是这几个，如果缺失，也是从这几个里面来
// [ ]
func firstMissingPositive(nums []int) int {
	n := len(nums)
	j := -1
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			j++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	j++
	for i := 0; i < j; i++ {
		source := nums[i]
		if source < 0 {
			source = -source
		}
		if source <= j && nums[source-1] > 0 {
			nums[source-1] = -nums[source-1]
		}
	}
	for i := 0; i < j; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return j + 1
}
