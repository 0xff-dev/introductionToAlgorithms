package leetcode

func sortedListToBST(head *ListNode) *TreeNode {
	return buildBST(head, nil)
}

func buildBST(start, end *ListNode) *TreeNode {
	if start == nil || start == end {
		return nil
	}
	one, two := start, start
	for two != end && two.Next != end {
		one = one.Next
		two = two.Next.Next
	}
	node := &TreeNode{
		Val: one.Val,
	}
	node.Left = buildBST(start, one)
	node.Right = buildBST(one.Next, end)
	return node
}
