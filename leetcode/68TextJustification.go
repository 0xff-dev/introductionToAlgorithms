package leetcode

import (
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	ans := make([]string, 0)
	l := maxWidth
	wc := 0
	// 必须有空格
	index := 0
	for ; index < len(words); index++ {
		w := words[index]
		//fmt.Printf("deal: %s, now ans: %v\n", w, ans)
		l -= len(w)
		wc++
		if l == 0 && wc == 1 {
			ans = append(ans, w)
			l, wc = maxWidth, 0
			continue
		}
		// a, b, c,d 5
		if l < wc-1 {
			// 剩余空个不够单词的兼具了, 这个单词就不可以选择,将前面的写入ans
			// 减少一个单词
			wc--
			// 开始写入单词的位置
			start := index - wc
			// 最后一个写入的单词的位置
			index--
			// 剩余多少个空格
			l += len(w)
			// 检查单词之间的间隔
			in := wc - 1
			sb := strings.Builder{}

			// 表示只有一个单词，直接写入，然后写入剩余的空格即可
			if in == 0 {
				sb.WriteString(words[index])
				for ; l > 0; l-- {
					sb.WriteByte(' ')
				}
			} else {
				// 按要求是尽可能平均，剩余尽量想做靠拢
				avg := l / in
				left := l % in
				for i, j := start, 0; i <= index; i, j = i+1, j+1 {
					sb.WriteString(words[i])
					if i != index {
						for a := avg; a > 0; a-- {
							sb.WriteByte(' ')
						}
						if j < left {
							sb.WriteByte(' ')
						}
					}

				}
			}
			ans = append(ans, sb.String())
			l, wc = maxWidth, 0
		}
	}
	if wc > 0 {
		start := index - wc
		index--
		in := wc - 1
		sb := strings.Builder{}

		if in == 0 {
			sb.WriteString(words[index])
			for ; l > 0; l-- {
				sb.WriteByte(' ')
			}
		} else {
			for i := start; i < index; i++ {
				sb.WriteString(words[i])
				sb.WriteByte(' ')
				l--
			}
			sb.WriteString(words[index])
			for ; l > 0; l-- {
				sb.WriteByte(' ')
			}
		}
		ans = append(ans, sb.String())
	}
	return ans
}
