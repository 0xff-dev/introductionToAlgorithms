package leetcode

import (
	"bytes"
)

func convert(s string, numRows int) string {
	strLen := len(s)
	if len(s) == 0 || strLen < numRows || numRows == 1 {
		return s
	}
	array := make([][]byte, numRows)
	for idx := range array {
		array[idx] = make([]byte, 0)
	}
	scrollIdx, dir := 0, 1
	for idx := 0; idx < strLen; idx++ {
		if scrollIdx == numRows-1 {
			dir = -1
		}
		if scrollIdx == 0 {
			dir = 1
		}
		array[scrollIdx] = append(array[scrollIdx], s[idx])
		scrollIdx += dir
	}
	buf := bytes.NewBuffer([]byte{})
	for _, arr := range array {
		buf.Write(arr)
	}
	return buf.String()
}

func convert2(s string, numRows int) string {
	buf := bytes.NewBufferString("")
	if numRows == 1 {
		return s
	}
	width := 2*(numRows-1)
	for idx := 0; idx < numRows; idx++ {
		step := 2*idx
		if idx == 0 || idx == numRows-1 {
			step = width
		}
		for walker := idx; walker < len(s); walker+=step {
			buf.WriteByte(s[walker])
			if step != width {
				step = width-step
			}
		}
	}

	return buf.String()
}