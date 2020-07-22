package offer

type ListNode struct {
	Val  int
	Next *ListNode
}

// 直接利用
func DeleteNode(node **ListNode) {
	*node = (*node).Next
}
