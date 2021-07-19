package leetcode

import "bytes"

var chars [26]int

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	r := make(map[string][]string, 0)
	for _, s := range strs {
		_s := coutSort(s)
		if _, ok := r[_s]; !ok {
			r[_s] = make([]string, 0)
		}

		r[_s] = append(r[_s], s)
	}
	for _, v := range r {
		res = append(res, v)
	}
	return res
}

func coutSort(s string) string {
	for _, c := range s {
		chars[c-'a']++
	}
	buf := bytes.NewBufferString("")
	for idx, c := range chars {
		for ; c > 0; c-- {
			buf.WriteByte(byte(idx + 'a'))
		}
		chars[idx] = 0
	}

	return buf.String()
}
