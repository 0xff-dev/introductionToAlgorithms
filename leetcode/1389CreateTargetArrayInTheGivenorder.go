package leetcode

func createTargetArray(nums []int, index []int) []int {
	ans := make([]int, len(index))
	for idx := 1; idx < len(index); idx++ {
		ans[idx] = -1
	}

	ans[0] = nums[0]
	for idx := 1; idx < len(index); idx++ {
		if ans[index[idx]] == -1 {
			ans[index[idx]] = nums[idx]
			continue
		}

		for w := idx - 1; w >= index[idx]; w-- {
			ans[w+1] = ans[w]
		}
		ans[index[idx]] = nums[idx]
	}
	return ans
}
