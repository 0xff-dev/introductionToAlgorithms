package leetcode

// 拓扑排序, 一共6个元素
// 0, 2, 1, 5, 3, 4
// 0, 1, 2, 3, 4, 5
func buildArray1920(nums []int) []int {
	ans := make([]int, len(nums))
	for idx, item := range nums {
		ans[idx] = nums[item]
	}
	return ans
}
