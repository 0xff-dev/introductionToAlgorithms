package leetcode

import "math/bits"

func bitwiseComplement(num int) int {
	if num == 0 {
		return 1
	}
	mask := (1 << bits.Len(uint(num))) - 1
	return num ^ mask
}
