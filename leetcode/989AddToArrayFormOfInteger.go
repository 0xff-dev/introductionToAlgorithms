package leetcode

import "fmt"

func addToArrayForm(num []int, k int) []int {
	ans := make([]int, 0)
	kForm := fmt.Sprintf("%d", k)
	l1, l2 := len(num), len(kForm)
	cf := 0
	i, j := l1-1, l2-1
	for i >= 0 || j >= 0 {
		t := cf
		if i >= 0 {
			t += num[i]
			i--
		}
		if j >= 0 {
			t += int(kForm[j] - '0')
			j--
		}
		cf = t / 10
		t %= 10
		ans = append(ans, t)
	}
	if cf > 0 {
		ans = append(ans, cf)
	}
	for s, e := 0, len(ans)-1; s < e; s, e = s+1, e-1 {
		ans[s], ans[e] = ans[e], ans[s]
	}
	return ans
}
