package leetcode

func countCommas3871(n int64) int64 {
	if n < 1000 {
		return 0
	}
	var (
		bits int64
		ret  int64
	)
	src := n
	for ; src > 0; src /= 10 {
		bits++
	}

	if bits < 7 {
		return n - 999
	}
	ret += 999000 // 999,000 * 1

	if bits < 10 {
		ret += (n - 999999) * 2
		return ret
	}
	ret += 999000000 * 2 // 999,000,000 * 2

	if bits < 13 {
		ret += (n - 999999999) * 3
		return ret
	}
	ret += 999000000000 * 3 // 999,000,000,000 * 3

	if bits < 16 {
		ret += (n - 999999999999) * 4
		return ret
	}
	ret += 999000000000000 * 4 // 999,000,000,000,000 * 4

	ret += (n - 999999999999999) * 5
	return ret
}
