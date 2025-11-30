package leetcode

func minSubarray(nums []int, p int) int {
	nums[0] %= p
	l := len(nums)
	for i := range l {
		nums[i] = (nums[i-1] + nums[i]) % p
	}
	if nums[l-1] == 0 {
		return 0
	}

	// 6, 3, 5, 2.  9
	// 6, 0, 5, 7
	ret := -1
	in := map[int]int{
		0: -1,
	}
	for i := range l {
		in[nums[i]] = i

		diff := (nums[i] - nums[l-1] + p) % p
		if index, ok := in[diff]; ok {
			if (ret == -1 || ret > i-index) && i-index != l {
				ret = i - index
			}
		}
	}
	return ret
}
