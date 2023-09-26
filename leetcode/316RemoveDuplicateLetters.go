package leetcode

func removeDuplicateLetters(s string) string {
	// 不像是动态规划
	// 不具备子问题的结构,具备
	// c a b,d,a
	// 暴力算法
	// 滑动窗口
	// stack
	// bcabc
	// a:1, b:2,c:2
	// abc
	count := [26]int{}
	visited := [26]bool{}
	for _, b := range []byte(s) {
		count[b-'a']++
	}
	stack := []byte{}
	for _, b := range []byte(s) {
		count[b-'a']--
		if visited[b-'a'] {
			continue
		}
		for len(stack) > 0 {
			l := len(stack)
			last := stack[l-1]
			if !(last > b && count[last-'a'] > 0) {
				break
			}
			visited[last-'a'] = false
			stack = stack[:l-1]
		}
		visited[b-'a'] = true
		stack = append(stack, b)
	}
	return string(stack)
}
