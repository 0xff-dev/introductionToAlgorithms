package leetcode

// 类似向左平铺，遇到不够高的，会降低，但是遇到超级高的会持续增强最高值
func minimizeArrayValue(nums []int) int {
	left, right := nums[0], 1000000000
	for left < right {
		mid := left + (right-left)/2
		buf := int64(0)
		i := 0
		for ; i < len(nums); i++ {
			buf += int64(mid - nums[i])
			if buf < 0 {
				break
			}
		}
		if i == len(nums) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
