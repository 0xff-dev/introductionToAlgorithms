package leetcode

import (
	"math"
)

const (
	maxStr8 = "2147483647"
	minStr8 = "2147483648"
)

func myAtoi(s string) int {
	bs := []byte(s)
	skip := byte(' ')
	index := 0
	neg := false
	repeatedSignedness := false
	start, end := -1, len(bs)
	for ; index < len(bs) && bs[index] == skip; index++ {
	}
	// 先过滤掉前导空格
	for ; index < len(bs); index++ {
		if bs[index] == '+' || bs[index] == '-' {
			if repeatedSignedness || start != -1 {
				end = index
				break
			}
			repeatedSignedness = true
			neg = bs[index] == '-'
			continue
		}
		if !(bs[index] >= '0' && bs[index] <= '9') {
			end = index
			break
		}
		if start == -1 {
			start = index
		}
	}
	if start == -1 {
		return 0
	}
	cut := bs[start:end]
	index = 0
	for ; index < len(cut) && cut[index] == '0'; index++ {
	}
	if index == len(cut) {
		return 0
	}
	cut = cut[index:]
	if len(cut) > len(maxStr8) {
		if neg {
			return math.MinInt32
		}
		return math.MaxInt32
	}

	if len(cut) == len(maxStr8) {
		if neg {
			// 看是是否比 minint32还要小
			for i := range len(minStr8) {
				if cut[i] == minStr8[i] {
					continue
				}
				if cut[i] < minStr8[i] {
					break

				}
				return math.MinInt32
			}
		} else {
			for i := range len(maxStr8) {
				if cut[i] == maxStr8[i] {
					continue
				}
				if cut[i] < maxStr8[i] {
					break
				}
				return math.MaxInt32
			}
		}
	}
	base := 0
	for i := 0; i < len(cut); i++ {
		base = base*10 + int(cut[i]-'0')
	}
	if neg {
		base = -base
	}
	return base
}
