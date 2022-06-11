package leetcode

func getConcatenation(nums []int) []int {
	length := len(nums)
	if length == 0 {
		return []int{}
	}
	ans := make([]int, length*2)
	for idx := 0; idx < length*2; idx++ {
		ans[idx] = nums[idx%length]
	}

	return ans
}
