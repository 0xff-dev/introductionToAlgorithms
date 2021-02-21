package leetcode

import (
	"sort"
)

// 感觉依次对比就可以，不需要做最长公共子序列
func findLongestWord(s string, d []string) string {
	sort.Slice(d, func(i, j int) bool {
		if len(d[i]) == len(d[j]) {
			return d[i] < d[j]
		}
		return len(d[i]) > len(d[j])
	})
	for _, str := range d {
		if len(str) > len(s) {
			continue
		}
		if isSubStr(s, str) {
			return str
		}
	}
	return ""
}

func isSubStr(a, b string) bool {
	length := len(b) + 1
	array := [2][]int{}
	for i := 0; i < 2; i++ {
		array[i] = make([]int, length)
	}
	idx := 1
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			tmp := 0
			if a[i-1] == b[j-1] {
				tmp = array[1-idx][j-1] + 1
			} else {
				tmp = max2(array[idx][j-1], array[1-idx][j])
			}
			array[idx][j] = tmp

			if tmp == length-1 {
				return true
			}
		}
		idx = 1 - idx
	}
	return false
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
