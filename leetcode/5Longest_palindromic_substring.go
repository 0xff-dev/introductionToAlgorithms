package leetcode

import (
	"bytes"
	"strings"
)

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	buf := bytes.NewBuffer([]byte{'_'})
	for _, b := range []byte(s) {
		buf.Write([]byte{b, '_'})
	}
	result := buf.Bytes()

	maxLen, index, step := 1, 0, 0
	for idx := 0; idx < len(result); idx++ {
		tmpStep := 1
		for idx-tmpStep >= 0 && idx+tmpStep < len(result) && result[idx-tmpStep] == result[idx+tmpStep] {
			tmpStep++
		}
		if 2*tmpStep-1 > maxLen {
			maxLen = 2*tmpStep - 1
			index = idx
			step = tmpStep
		}
	}
	step--
	r := result[index-step : index+step]
	buffer := bytes.NewBufferString("")
	for _, s := range strings.Split(string(r), "_") {
		buffer.WriteString(s)
	}
	return buffer.String()
}
