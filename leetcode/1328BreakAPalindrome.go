package leetcode

func breakPalindrome(palindrome string) string {
	l := len(palindrome)
	if l == 0 || l == 1 {
		return ""
	}
	source := []byte(palindrome)
	mid := l / 2
	i := 0

	for ; i < l; i++ {
		if l&1 == 1 && i == mid {
			continue
		}
		if source[i] != 'a' {
			break
		}
	}
	set := byte('a')
	if i == l {
		i--
		set = 'b'
	}
	source[i] = set
	return string(source)
}
