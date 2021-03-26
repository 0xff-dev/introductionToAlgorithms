package leetcode

func maxProduct(nums []int) int {
	r := make([]int, len(nums))
	r[0] = nums[0]
	maxRes := r[0]
	for idx := 1; idx < len(nums); idx++ {
		if r[idx-1] == 0 {
			r[idx] = nums[idx]
		} else {
			r[idx] = r[idx-1] * nums[idx]
		}
		if r[idx] > maxRes {
			maxRes = r[idx]
		}

		if r[idx] != 0 {
			for i := idx-1; i >= 0 && r[i] != 0; i-- {
				tmp := r[idx] / r[i]
				if tmp > maxRes {
					maxRes = tmp
				}
			}
		}
	}
	return maxRes
}
