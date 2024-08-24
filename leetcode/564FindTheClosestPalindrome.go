package leetcode

import (
	"strconv"
)

// Convert to palindrome keeping first half constant.
func convert564(num int64) int64 {
	s := strconv.FormatInt(num, 10)
	n := len(s)
	l := (n - 1) / 2
	r := n / 2
	chars := []rune(s)
	for l >= 0 {
		chars[r] = chars[l]
		r++
		l--
	}
	newS := string(chars)
	result, _ := strconv.ParseInt(newS, 10, 64)
	return result
}

// Find the next palindrome, just greater than n.
func nextPalindrome(num int64) int64 {
	left := int64(0)
	right := num
	var ans int64
	for left <= right {
		mid := (right-left)/2 + left
		palin := convert564(mid)
		if palin < num {
			ans = palin
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

// Find the previous palindrome, just smaller than n.
func previousPalindrome(num int64) int64 {
	left := num
	right := int64(1e18)
	var ans int64
	for left <= right {
		mid := (right-left)/2 + left
		palin := convert564(mid)
		if palin > num {
			ans = palin
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

func nearestPalindromic(n string) string {
	num, _ := strconv.ParseInt(n, 10, 64)
	a := nextPalindrome(num)
	b := previousPalindrome(num)
	if diff(a, num) <= diff(b, num) {
		return strconv.FormatInt(a, 10)
	}
	return strconv.FormatInt(b, 10)
}

func diff(a, b int64) int64 {
	if a > b {
		return a - b
	}
	return b - a
}
