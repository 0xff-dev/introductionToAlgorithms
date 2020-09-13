package challengeProgramming

import "bytes"

/*
	给一个字符串，每次从该字符串的首部，尾部娶一个字符，目标是构建一个字典序最小的字符串T
*/

func LexicalOrder(str string) string {
	s, e := 0, len(str)-1
	if len(str) == 0 {
		return ""
	}

	buf := bytes.NewBufferString("")
	for s <= e {
		appendLeft := true
		for index := 0; s+index <= e; index++ {
			if str[s+index] < str[e-index] {
				appendLeft = true
				break
			} else if str[s+index] > str[e-index] {
				appendLeft = false
				break
			}
		}
		if appendLeft {
			buf.WriteByte(str[s])
			s++
		} else {
			buf.WriteByte(str[e])
			e--
		}
	}
	return buf.String()
}
