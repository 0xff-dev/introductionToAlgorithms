package leetcode

import "strconv"

const (
	minStr = "2147483648"
	maxStr = "2147483647"
)

func reverse7(x int) int {
	minL, maxL := len(minStr), len(maxStr)
	neg := x < 0
	intStr := strconv.Itoa(x)
	if neg {
		intStr = intStr[1:]
		if len(intStr) > minL {
			return 0
		}
		if len(intStr) == minL {
			i, j := 0, minL-1
			for ; i < minL; i, j = i+1, j-1 {
				if minStr[i] == intStr[j] {
					continue
				}
				if minStr[i] > intStr[j] {
					break
				}
				return 0
			}
		}
		base := 0
		for i := len(intStr) - 1; i >= 0; i-- {
			base = base*10 + int(intStr[i]-'0')
		}
		return -base
	}
	if len(intStr) > maxL {
		return 0
	}
	if len(intStr) == maxL {
		i, j := 0, maxL-1
		for ; i < maxL; i, j = i+1, j-1 {
			if maxStr[i] == intStr[j] {
				continue
			}
			if maxStr[i] > intStr[j] {
				break
			}
			return 0
		}
	}
	base := 0
	for i := len(intStr) - 1; i >= 0; i-- {
		base = base*10 + int(intStr[i]-'0')
	}
	return base
}
