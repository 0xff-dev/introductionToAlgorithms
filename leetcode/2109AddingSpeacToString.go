package leetcode

import "strings"

func addSpaces(s string, spaces []int) string {
	sb := strings.Builder{}

	spaceIdx := 0
	for i, b := range s {
		if spaceIdx < len(spaces) && i == spaces[spaceIdx] {
			sb.WriteByte(' ')
			spaceIdx++
		}
		sb.WriteByte(byte(b))
	}

	return sb.String()
}
