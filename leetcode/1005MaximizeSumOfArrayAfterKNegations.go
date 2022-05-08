package leetcode

import "sort"

func _abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}
func largestSumAfterKNegations(nums []int, k int) int {
	// 如果负数比k小，
	negations := 0
	hasZero := false
	asbMin := 101
	ans := 0
	for _, n := range nums {
		if n == 0 {
			hasZero = true
			continue
		}
		abs := _abs(n)
		ans += abs
		if abs < asbMin {
			asbMin = abs
		}
		if n < 0 {
			negations++
		}
	}
	if negations == k {
		return ans
	}
	sort.Ints(nums)
	if negations < k {
		if hasZero {
			return ans
		}
		if (k-negations)&1 == 1 {
			ans -= 2 * asbMin
		}
		return ans
	}
	for idx, n := range nums {
		if idx < k {
			continue
		}
		if n > 0 {
			break
		}
		ans += 2 * n
	}

	return ans
}
