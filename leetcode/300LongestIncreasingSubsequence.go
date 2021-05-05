package leetcode

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	r := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		r[i] = 1
	}
	maxLength := 1
	for idx := 1; idx < len(nums); idx++ {
		for j := 0; j < idx; j++ {
			if nums[j] < nums[idx] && r[j]+1 > r[idx] {
				r[idx] = r[j] + 1
			}
		}
		if r[idx] > maxLength {
			maxLength = r[idx]
		}
	}
	return maxLength
}
