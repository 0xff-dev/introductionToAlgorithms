package leetcode

func isBadVersion(n int) bool {
	return false
}

func firstBadVersion(n int) int {
	l, r := 1, n+1
	ans := 1
	for l < r {
		mid := l + (r-l)/2
		if isBadVersion(mid) {
			r = mid
			ans = mid
			continue
		}
		l = mid + 1
	}
	return ans
}
