package leetcode

func rotate1(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}

	rotateReverse(nums)
	rotateReverse(nums[:k])
	rotateReverse(nums[k:])
}

func rotateReverse(nums []int) {
	for s, e := 0, len(nums)-1; s < e; s, e = s+1, e-1 {
		nums[s], nums[e] = nums[e], nums[s]
	}
}
