package leetcode

func longestAlternatingSubarray(nums []int, threshold int) int {
	size := len(nums)
	start, bit := -1, -1
	var ret, tmp int
	for index := range size {
		tmp = nums[index] & 1
		if start != -1 {
			if bit == tmp || nums[index] > threshold {
				ret = max(ret, index-start)
				start = -1
				if tmp == 0 && nums[index] <= threshold {
					start = index
					bit = 0
				}
				continue
			}
			bit = tmp
			continue
		}

		if tmp == 0 && nums[index] <= threshold {
			start = index
			bit = 0
		}
	}
	if start != -1 {
		ret = max(ret, size-start)
	}
	return ret
}
