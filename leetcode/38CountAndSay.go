package leetcode

import (
	"bytes"
	"fmt"
)

func countAndSay(n int) string {
	ans := "1"
	if n == 1 {
		return ans
	}

	buf := bytes.NewBufferString("")
	for loop := 2; loop <= n; loop++ {
		count := 0
		for i := 0; i < len(ans); i++ {
			if i == 0 || ans[i] == ans[i-1] {
				count++
				continue
			}
			buf.WriteString(fmt.Sprintf("%d%c", count, ans[i-1]))
			count = 1
		}
		buf.WriteString(fmt.Sprintf("%d%c", count, ans[len(ans)-1]))
		ans = buf.String()
		buf.Reset()
	}
	return ans
}
