package leetcode

import (
	"strings"
)

// 这个也是硬比较问题，一般涉及到多个low letter的比较问题，都要考虑trie结构
// 本周末写一下前缀树，重写820，以及745的问题
func minimumLengthEncoding(words []string) int {
	length := len(words)
	if length == 0 {
		return 0
	}
	//sb := strings.Builder{}
	//now := words[0]
	//for idx := 1; idx < length; idx++ {
	//	if strings.HasSuffix(now, words[idx]) {
	//		continue
	//	}
	//	if strings.HasSuffix(words[idx], now) {
	//		now = words[idx]
	//		continue
	//	}
	//
	//	sb.WriteString(now)
	//	sb.WriteByte('#')
	//	now = words[idx]
	//}
	//sb.WriteString(now)
	//s := sb.String()
	//fmt.Println("ans: ", s)
	//return len(s) + 1
	check := make([]uint8, len(words))
	// pa dpaa paa a
	// 1        0
	// time, me, bell
	//        1
	for idx := 0; idx < length-1; idx++ {
		if check[idx] == 1 {
			continue
		}
		now := words[idx]
		for in := idx + 1; in < length; in++ {
			if strings.HasSuffix(now, words[in]) {
				check[in] = 1
				continue
			}
			if strings.HasSuffix(words[in], now) {
				check[idx] = 1
				now = words[idx]
				continue
			}
		}
	}
	ans := 0

	for idx := 0; idx < length; idx++ {
		if check[idx] == 0 {
			ans += len(words[idx]) + 1
		}
	}
	return ans
}
