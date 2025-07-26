package leetcode

import "math"

func maxSubarrays(n int, conflictingPairs [][]int) int64 {
	bMin1 := make([]int, n+1)
	bMin2 := make([]int, n+1)
	for i := 0; i <= n; i++ {
		bMin1[i] = math.MaxInt32
		bMin2[i] = math.MaxInt32
	}
	for _, pair := range conflictingPairs {
		a := min(pair[0], pair[1])
		b := max(pair[0], pair[1])
		if bMin1[a] > b {
			bMin2[a] = bMin1[a]
			bMin1[a] = b
		} else if bMin2[a] > b {
			bMin2[a] = b
		}
	}
	res, ib1, b2 := int64(0), n, math.MaxInt32
	delCount := make([]int64, n+1)
	for i := n; i >= 1; i-- {
		if bMin1[ib1] > bMin1[i] {
			b2 = min(b2, bMin1[ib1])
			ib1 = i
		} else {
			b2 = min(b2, bMin1[i])
		}
		res += int64(min(bMin1[ib1], n+1) - i)
		delCount[ib1] += int64(min(min(b2, bMin2[ib1]), n+1) - min(bMin1[ib1], n+1))
	}
	maxVal := int64(0)
	for _, v := range delCount {
		maxVal = max(maxVal, v)
	}
	return res + maxVal
}
