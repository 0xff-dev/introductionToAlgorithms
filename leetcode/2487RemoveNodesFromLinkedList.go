package leetcode

func removeNodes(head *ListNode) *ListNode {
	// 5, 2, 13, 3, 8
	// 13 8
	stack := []*ListNode{head}
	w := head.Next
	for ; w != nil; w = w.Next {
		for len(stack) > 0 && stack[len(stack)-1].Val < w.Val {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, w)
	}
	for i := 1; i < len(stack); i++ {
		stack[i-1].Next = stack[i]
	}
	return stack[0]
}
