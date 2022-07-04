package leetcode

func candy(ratings []int) int {
	length := len(ratings)
	left := make([]int, length)
	right := make([]int, length)

	left[0], right[length-1] = 1, 1
	for idx := 1; idx < length; idx++ {
		left[idx] = 1
		if ratings[idx] > ratings[idx-1] {
			left[idx] = left[idx-1] + 1
		}
	}
	ans := left[length-1]
	if right[length-1] > ans {
		ans = right[length-1]
	}
	for idx := length - 2; idx >= 0; idx-- {
		right[idx] = 1
		if ratings[idx] > ratings[idx+1] {
			right[idx] = right[idx+1] + 1
		}

		less := left[idx]
		if less < right[idx] {
			less = right[idx]
		}
		ans += less
	}
	return ans
}
