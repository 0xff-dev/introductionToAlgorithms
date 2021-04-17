package leetcode

import "fmt"

func partition1(s string) [][]string {
	r := make([][]string, 0)
	palineAux(s, 0, []string{}, &r)
	return r
}

func palineAux(s string, index int, now []string, r *[][]string) {
	if index >= len(s) {
		dest := make([]string, len(now))
		copy(dest, now)
		*r = append(*r, dest)
		return
	}

	for i := index; i < len(s); i++ {
		tmpStr := s[index : i+1]
		if isPalinfrome(tmpStr) {
			fmt.Println("now str: ", tmpStr)
			palineAux(s, i+1, append(now, tmpStr), r)
		}
	}
}

func isPalinfrome(str string) bool {
	for s, e := 0, len(str)-1; s < e; s, e = s+1, e-1 {
		if str[s] != str[e] {
			return false
		}
	}
	return true
}
