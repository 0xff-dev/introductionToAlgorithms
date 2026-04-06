package leetcode

import "sort"

func gcd3867(a, b int) int {
	var mod int
	for b != 0 {
		mod = a % b
		a, b = b, mod
	}
	return a
}

func gcdSum(nums []int) int64 {
	var ret int64
	prefix := make([]int, len(nums))
	mxi := 0
	for i := range nums {
		mxi = max(mxi, nums[i])
		prefix[i] = gcd3867(mxi, nums[i])
	}
	sort.Ints(prefix)

	for s, e := 0, len(prefix)-1; s < e; s, e = s+1, e-1 {
		ret += int64(gcd3867(prefix[s], prefix[e]))
	}
	return ret
}
