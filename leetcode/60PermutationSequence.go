package leetcode

import (
	"fmt"
)

func getPermutation(n int, k int) string {
	factorial := make([]int, n+1)
	factorial[0] = 1
	remainingNums := make([]int, n+1)
	for idx := 1; idx <= n; idx++ {
		factorial[idx] = idx * factorial[idx-1]
		remainingNums[idx] = idx
	}

	base, remaining, count := 0, k, n
	used := make(map[int]struct{})

	for remaining > 1 {
		baseIdx := remaining / factorial[count-1]
		remaining %= factorial[count-1]
		modEqualZero := true
		if remaining != 0 {
			baseIdx++
			modEqualZero = false
		}
		base = base*10 + remainingNums[baseIdx]
		used[remainingNums[baseIdx]] = struct{}{}
		remainingNums = []int{0}
		if modEqualZero {
			for idx := n; idx >= 1; idx-- {
				if _, ok := used[idx]; !ok {
					remainingNums = append(remainingNums, idx)
				}
			}
			break
		}

		for idx := 1; idx <= n; idx++ {
			if _, ok := used[idx]; !ok {
				remainingNums = append(remainingNums, idx)
			}
		}
		count--
	}
	for idx := 1; idx < len(remainingNums); idx++ {
		base = base*10 + remainingNums[idx]
	}
	return fmt.Sprintf("%d", base)
}
