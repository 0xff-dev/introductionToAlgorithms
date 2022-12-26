package leetcode

import "strings"

func generateTheString(n int) string {
	sb := strings.Builder{}
	even := n&1 == 0
	c := byte('a')
	if even {
		sb.WriteByte(c)
		c++
		n--
	}

	for i := 0; i < n; i++ {
		sb.WriteByte(c)
	}
	return sb.String()
}
