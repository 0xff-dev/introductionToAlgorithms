package leetcode

func nextLargerNodes(head *ListNode) []int {
	monotonicStack := make([]int, 0)
	for walker := head; walker != nil; walker = walker.Next {
		monotonicStack = append(monotonicStack, walker.Val)
	}
	l := len(monotonicStack)
	ans := make([]int, l)
	ans[l-1] = 0
	stackIndex := l - 1
	for idx := l - 2; idx >= 0; idx-- {
		for stackIndex < l && monotonicStack[stackIndex] <= monotonicStack[idx] {
			stackIndex++
		}
		if stackIndex == l {
			ans[idx] = 0
			stackIndex = l - 1
			monotonicStack[stackIndex] = monotonicStack[idx]
			continue
		}
		ans[idx] = monotonicStack[stackIndex]
		stackIndex--
		monotonicStack[stackIndex] = monotonicStack[idx]
	}
	return ans
}
