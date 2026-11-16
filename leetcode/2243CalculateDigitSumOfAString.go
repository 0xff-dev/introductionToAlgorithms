package leetcode

import (
	"strconv"
	"strings"
)

func digitSum(s string, k int) string {
	var buf strings.Builder

	replace := func(src string) string {
		buf.Reset()
		base := 0
		for i := range src {
			if i%k == 0 {
				if i != 0 {
					_new := strconv.Itoa(base)
					buf.WriteString(_new)
				}
				base = 0
			}
			base += int(src[i] - '0')
		}
		_new := strconv.Itoa(base)
		buf.WriteString(_new)
		return buf.String()
	}

	for len(s) > k {
		s = replace(s)
	}
	return s
}
