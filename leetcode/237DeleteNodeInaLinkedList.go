package leetcode

func deleteNode1(node *ListNode) {
	*node = *node.Next
}
