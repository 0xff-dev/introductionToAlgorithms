package leetcode

import "strings"

func lexPalindromicPermutation3734(s string, target string) string {
	n := len(s)
	// Special case: length of 1
	if n == 1 {
		if s > target {
			return s
		}
		return ""
	}

	// Count the frequency of each character
	cnt := make([]int, 26)
	for _, c := range s {
		cnt[c-'a']++
	}

	// Check if it can form a palindrome and record the characters with odd occurrences
	oddChar := ""
	for i := 0; i < 26; i++ {
		if cnt[i]%2 == 1 {
			// More than one character appears an odd number of times, cannot form a palindrome
			if oddChar != "" {
				return ""
			}
			oddChar = string(rune('a' + i))
		}
		cnt[i] /= 2 // It takes only half the characters to construct the left half
	}

	prefix := ""

	check := func(c byte) bool {
		left := prefix + string(c)
		for i := 25; i >= 0; i-- {
			left += strings.Repeat(string(rune('a'+i)), cnt[i])
		}

		// reverse left
		reversedLeft := reverseString3734(left)
		palindrome := left + oddChar + reversedLeft

		return palindrome > target
	}

	// Construct the left part of each digit greedily
	for i := 0; i < n/2; i++ {
		found := false
		// Try to place the smallest character in lexicographical order
		for j := 0; j < 26; j++ {
			if cnt[j] == 0 {
				continue
			}

			cnt[j]--
			if check(byte('a' + j)) {
				// If the constructed palindrome is greater than target, choose the character
				prefix += string(byte('a' + j))
				found = true
				break
			} else {
				cnt[j]++ // Not meeting the conditions, reset the counter
			}
		}
		if !found {
			return "" // Cannot construct a palindrome larger than target
		}

		if prefix[i] > target[i] { // prefix is already greater than target
			left := prefix
			for j := 0; j < 26; j++ {
				left += strings.Repeat(string(rune('a'+j)), cnt[j])
			}
			palindrome := left + oddChar + reverseString3734(left)
			return palindrome
		}
	}

	// Construct the final palindrome string
	ans := prefix + oddChar + reverseString3734(prefix)
	return ans
}

func reverseString3734(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
