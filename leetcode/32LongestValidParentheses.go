package leetcode

func longestValidParentheses(s string) int {
	bs := []byte(s)
	ret, i, add := 0, 0, 0
	size := len(bs)

	stack := make([]int, size+1)
	top := 0

	for i < size {
		if bs[i] == '(' {
			top++
			stack[top] = 0
			i++
			continue
		}
		//    ) ()()
		//2 遇到 ), 无法继续连续
		if top == 0 {
			ret = max(ret, stack[0])
			stack[0] = 0
			i++
			continue
		}
		for top > 0 && i < size && bs[i] == ')' {
			add = stack[top] + 2
			top--
			stack[top] += add
			i++
		}
		ret = max(ret, stack[top])
	}
	return ret
}
