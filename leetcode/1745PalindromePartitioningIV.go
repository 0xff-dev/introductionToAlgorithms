package leetcode

func isPalindrome1745(s string) bool {
	l, r := 0, len(s)-1
	for ; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}

func checkPartitioning(s string) bool {
	l := len(s)
	if l == 3 {
		return true
	}
	// dp[i] 表示s[i:] 是否可以分割成两个回文串
	left := make([]bool, l)
	left[0] = true
	for i := 1; i < l; i++ {
		left[i] = isPalindrome1745(s[:i+1])
	}
	right := make([]bool, l)
	right[l-1] = true
	for i := l - 2; i >= 0; i-- {
		right[i] = isPalindrome1745(s[i:])
	}

	for i := 1; i < l-1; i++ {
		for ll, r := i-1, i+1; ll >= 0 && r <= l-1; ll, r = ll-1, r+1 {
			if left[ll] && right[r] {
				return true
			}
			if s[ll] != s[r] {
				break
			}
		}
	}
	for i := 1; i < l-2; i++ {
		if s[i] != s[i+1] {
			continue
		}
		for ll, r := i-1, i+2; ll >= 0 && r <= l-1; ll, r = ll-1, r+1 {
			if left[ll] && right[r] {
				return true
			}
			if s[ll] != s[r] {
				break
			}
		}
	}
	return false
}
