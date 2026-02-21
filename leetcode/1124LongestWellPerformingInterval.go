package leetcode

func longestWPI(hours []int) int {
	n := len(hours)
	score := make([]int, n)

	for i, h := range hours {
		if h > 8 {
			score[i] = 1
		} else {
			score[i] = -1
		}
	}

	// 构造前缀和数组
	diff := make([]int, n+1)
	for i := 1; i <= n; i++ {
		diff[i] = diff[i-1] + score[i-1]
	}

	// 构建单调递减栈，存储diff的索引
	stack := []int{}
	for i := 0; i <= n; i++ {
		if len(stack) == 0 || diff[i] < diff[stack[len(stack)-1]] {
			stack = append(stack, i)
		}
	}

	maxLength := 0
	// 从后向前遍历diff数组，试图找到最长满足条件的区间
	for i := n; i >= 0; i-- {
		for len(stack) > 0 && diff[i] > diff[stack[len(stack)-1]] {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if i-j > maxLength {
				maxLength = i - j
			}
		}
	}

	return maxLength
}
