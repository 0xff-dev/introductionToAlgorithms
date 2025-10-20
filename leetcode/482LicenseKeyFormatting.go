package leetcode

import (
	"bytes"
	"slices"
	"strings"
)

func isAlphanumericASCII(b byte) bool {
	return (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
}

func licenseKeyFormatting(s string, k int) string {
	s = strings.ToUpper(s)
	l := len(s)
	bs := []byte(s)
	buf := bytes.NewBuffer([]byte{})
	count := 0
	for i := l - 1; i >= 0; i-- {
		if isAlphanumericASCII(bs[i]) {
			count++
			buf.WriteByte(bs[i])
			if count == k {
				buf.WriteByte('-')
				count = 0
			}
		}
	}
	ret := buf.Bytes()
	if len(ret) == 0 {
		return ""
	}
	slices.Reverse(ret)
	if ret[0] == '-' {
		ret = ret[1:]
	}
	return string(ret)
}
