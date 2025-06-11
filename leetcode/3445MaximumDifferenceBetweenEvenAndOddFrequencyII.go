package leetcode

import "math"

func maxDifference3445(s string, k int) int {
	n := len(s)
	ans := math.MinInt32
	for _, a := range []byte{'0', '1', '2', '3', '4'} {
		for _, b := range []byte{'0', '1', '2', '3', '4'} {
			if a == b {
				continue
			}
			best := make([]int, 4)
			for i := range best {
				best[i] = math.MaxInt32
			}
			cntA, cntB := 0, 0
			prevA, prevB := 0, 0
			left := -1
			for right := 0; right < n; right++ {
				if s[right] == a {
					cntA++
				}
				if s[right] == b {
					cntB++
				}
				for right-left >= k && cntB-prevB >= 2 {
					leftStatus := getStatus(prevA, prevB)
					if prevA-prevB < best[leftStatus] {
						best[leftStatus] = prevA - prevB
					}
					left++
					if s[left] == a {
						prevA++
					}
					if s[left] == b {
						prevB++
					}
				}
				rightStatus := getStatus(cntA, cntB)
				if best[rightStatus^0b10] != math.MaxInt32 {
					current := (cntA - cntB) - best[rightStatus^0b10]
					if current > ans {
						ans = current
					}
				}
			}
		}
	}
	return ans
}

func getStatus(cntA, cntB int) int {
	return ((cntA & 1) << 1) | (cntB & 1)
}
