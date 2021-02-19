package leetcode

import (
	"bytes"
	"fmt"
	"sort"
)

func largestNumber(nums []int) string {
	strs := toString(nums)
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})
	r := bytes.NewBufferString("")
	for _, s := range strs {
		r.WriteString(s)
	}
	str := r.String()
	i := 0
	for ; i < len(str) && str[i] == '0'; i++ {
	}
	if i == len(str) {
		return "0"
	}
	return str[i:]
}

func toString(nums []int) []string {
	r := make([]string, 0)
	for _, n := range nums {
		r = append(r, fmt.Sprintf("%d", n))
	}
	return r
}
