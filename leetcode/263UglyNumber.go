package leetcode

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for {
		if n%2 != 0 {
			break
		}
		n /= 2
	}
	for {
		if n%3 != 0 {
			break
		}
		n /= 3
	}
	for {
		if n%5 != 0 {
			break
		}
		n /= 5
	}
	return n == 1
}
