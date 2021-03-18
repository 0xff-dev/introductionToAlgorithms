package leetcode

func titleToNumber(columnTitle string) int {
	base26, sum := 1, 0
	const at = '@'

	for idx := len(columnTitle) - 1; idx >= 0; idx-- {
		sum += base26 * int(columnTitle[idx]-at)
		base26 *= 26
	}
	return sum
}
