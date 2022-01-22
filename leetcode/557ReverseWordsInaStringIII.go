package leetcode

import "bytes"

// Let's go
// og s'teL
func reverseWords3(s string) string {
	buf := bytes.NewBufferString("")
	stack := make([]string, 0)
	for idx := len(s) - 1; idx >= 0; idx-- {
		if s[idx] == ' ' {
			stack = append(stack, buf.String())
			buf.Reset()
			continue
		}
		buf.WriteByte(s[idx])
	}
	stack = append(stack, buf.String())
	buf.Reset()
	first := true
	for idx := len(stack) - 1; idx >= 0; idx-- {
		writeObj := stack[idx]
		if !first {
			writeObj = " " + writeObj
		}
		first = false
		buf.WriteString(writeObj)
	}
	// Let's go
	// og s'teL
	return buf.String()
}
