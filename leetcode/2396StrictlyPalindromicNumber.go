package leetcode

func helper2396(n, base int) bool {
	store := make([]int, 0)
	// 3/2
	// 4 2
	// 0 0
	// 1
	for n >= base {
		store = append(store, n%base)
		n /= base
	}
	store = append(store, n)
	for s, e := 0, len(store); s < e; s, e = s+1, e-1 {
		if store[s] != store[e] {
			return false
		}
	}
	return true
}
func isStrictlyPalindromic(n int) bool {
	for i := 2; i <= n-2; i++ {
		if !helper2396(n, i) {
			return false
		}
	}
	return true
}
