package leetcode

import (
	"sort"
	"strings"
)

type tmp1451 struct {
	i int
	s string
}

func arrangeWords(text string) string {
	if len(text) == 0 {
		return ""
	}
	bs := strings.Split(text, " ")
	bs[0] = strings.ToLower(bs[0])

	indies := make([]tmp1451, len(bs))
	for i := 0; i < len(bs); i++ {
		indies[i] = tmp1451{
			i: i, s: bs[i],
		}
	}
	sort.Slice(indies, func(i, j int) bool {
		a, b := len(indies[i].s), len(indies[j].s)
		if a == b {
			return indies[i].i < indies[j].i
		}
		return a < b
	})
	ans := make([]string, len(bs))
	for i := 0; i < len(bs); i++ {
		ans[i] = indies[i].s
	}
	first := []byte(ans[0])
	first[0] -= 32
	ans[0] = string(first)
	return strings.Join(ans, " ")
}
