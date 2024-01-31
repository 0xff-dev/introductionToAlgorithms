package leetcode

/*
func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	// 用下表跳?
	for idx := 0; idx < len(temperatures)-1; idx++ {
		next := idx + 1
		for ; next < len(temperatures) && temperatures[next] <= temperatures[idx]; next++ {
		}
		if next < len(temperatures) {
			ans[idx] = next - idx
		}
	}

	return ans
}
*/

func dailyTemperatures(temperatures []int) []int {
	l := len(temperatures)
	ans := make([]int, l)
	stack := make([]int, l)
	stack[l-1] = l - 1 //指向最后一个元素
	cur := l - 1
	for idx := l - 2; idx >= 0; idx-- {
		for ; cur < l && temperatures[stack[cur]] <= temperatures[idx]; cur++ {
		}
		if cur != l {
			ans[idx] = stack[cur] - idx
		}
		cur--
		stack[cur] = idx
	}
	return ans
}
