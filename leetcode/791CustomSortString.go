package leetcode

import (
	"bytes"
	"sort"
)

func customSortString(order string, str string) string {
	orderMap := make([]uint8, 26)
	ordered := uint8(1)
	for _, b := range []byte(order) {
		orderMap[b-'a'] = ordered
		ordered++
	}

	strBs := []byte(str)
	sort.Slice(strBs, func(i, j int) bool {
		return orderMap[strBs[i]-'a'] < orderMap[strBs[j]-'a']
	})

	return string(strBs)
}

func customSortString1(order string, str string) string {
	buf := bytes.NewBufferString("")
	chars := [26]uint8{}
	for _, c := range str {
		chars[c-'a']++
	}

	for _, c := range order {
		if count := chars[c-'a']; count > 0 {
			for ; count > 0; count-- {
				buf.WriteByte(byte(c))
			}
			chars[c-'a'] = 0
		}
	}

	for idx, c := range chars {
		for ; c > 0; c-- {
			buf.WriteByte(byte(idx) + 'a')
		}
	}

	return buf.String()
}
