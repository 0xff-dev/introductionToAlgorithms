package leetcode

import (
	"bytes"
)

func longestDiverseString(a int, b int, c int) string {
	if a == b && b == c && c == 0 {
		return ""
	}

	store := map[byte]int{
		'a': a,
		'b': b,
		'c': c,
	}

	buf := bytes.NewBufferString("")
	first, second := byte(' '), byte(' ')
	fillChar, count := maxChar(a, b, c)
	if count == 0 {
		return ""
	}

	buf.WriteByte(fillChar)
	first, second = second, fillChar
	count--

	second = fillChar
	if count >= 1 {
		first, second = second, fillChar
		buf.WriteByte(fillChar)
		count--
	}
	store[fillChar] = count

	for {
		// 填充了同样的数据 aa, ' 'b
		_max := -1
		tmpFill := fillChar
		if first == second {
			for k, v := range store {
				if k == tmpFill {
					continue
				}
				if v > 0 && v > _max {
					_max = v
					count = _max
					fillChar = k
				}
			}
		} else {
			for k, v := range store {
				if v > 0 && v > _max {
					_max = v
					count = v
					fillChar = k
				}
			}
		}

		if _max == -1 || (first == second && fillChar == second) {
			break
		}

		buf.WriteByte(fillChar)
		tmpSecond := second
		flag := first == second
		first, second = second, fillChar
		count--

		if !flag && fillChar != tmpSecond && count >= 1 {
			buf.WriteByte(fillChar)
			first, second = second, fillChar
			count--
		}
		store[fillChar] = count
	}

	return buf.String()
}

func maxChar(a, b, c int) (byte, int) {
	char := byte('a')
	max := a
	if b > max {
		char = 'b'
		max = b
	}
	if c > max {
		char = 'c'
		max = c
	}
	return char, max
}
