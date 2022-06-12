package leetcode

func maximumUniqueSubarray(nums []int) int {
	length := len(nums)
	ans := 0
	if length == 0 {
		return ans
	}

	check := make(map[int]int)
	start, end, sum := 0, 0, 0
	for ; end < length; end++ {
		sum += nums[end]
		if repeatIdx, ok := check[nums[end]]; ok {
			for idx := start; idx <= repeatIdx; idx++ {
				sum -= nums[idx]
				delete(check, nums[idx])
			}
			start = repeatIdx + 1
		}
		check[nums[end]] = end
		if sum > ans {
			ans = sum
		}
	}
	return ans
}
