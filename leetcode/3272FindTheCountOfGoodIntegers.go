package leetcode

import (
	"sort"
	"strconv"
	"strings"
)

func countGoodIntegers(n int, k int) int64 {
	dict := make(map[string]bool)
	base := intPow(10, (n-1)/2)
	skip := n & 1
	/* Enumerate the number of palindrome numbers of n digits */
	for i := base; i < base*10; i++ {
		s := strconv.Itoa(i)
		rev := reverseString3272(s)
		s += rev[skip:]
		palindromicInteger, _ := strconv.ParseInt(s, 10, 64)
		/* If the current palindrome number is a k-palindromic integer */
		if palindromicInteger%int64(k) == 0 {
			chars := strings.Split(s, "")
			sort.Strings(chars)
			dict[strings.Join(chars, "")] = true
		}
	}

	factorial := make([]int64, n+1)
	factorial[0] = 1
	for i := 1; i <= n; i++ {
		factorial[i] = factorial[i-1] * int64(i)
	}

	var ans int64 = 0
	for s := range dict {
		cnt := make([]int, 10)
		for _, c := range s {
			cnt[c-'0']++
		}
		/* Calculate permutations and combinations */
		tot := int64(n-cnt[0]) * factorial[n-1]
		for _, x := range cnt {
			tot /= factorial[x]
		}
		ans += tot
	}

	return ans
}

func intPow(a, b int) int {
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}

func reverseString3272(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
