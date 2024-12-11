package leetcode

import "sort"

func maximumBeauty2779(nums []int, k int) int {
	sort.Ints(nums)
	ans := 0
	start, end := 0, 0
	// 排序以后，能够出现交集的情况就是
	// 最小的元素定义了右边界
	// 最大的元素定义了左边界
	// re 就是交集的右边界
	re := nums[0] + k

	for ; end < len(nums); end++ {
		// 还在同一个范围
		s := nums[end] - k
		// 只要当前的s还是比re小，就可以继续增加
		if s <= re {
			continue
		}
		diff := end - start
		ans = max(ans, diff)

		// 开始减少左侧的数据
		// 尝试移除start
		start++
		for ; start <= end; start++ {
			re = nums[start] + k
			if s <= re {
				break
			}
		}
	}
	diff := end - start
	ans = max(ans, diff)

	return ans
}
