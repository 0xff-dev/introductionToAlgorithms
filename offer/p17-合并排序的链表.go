package offer

// 尝试递归？
func MergeSortList(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var node *ListNode
	if l1.Val > l2.Val {
		node = &ListNode{l2.Val, nil}
		node.Next = MergeSortList(l1, l2.Next)
	} else {
		node = &ListNode{l1.Val, nil}
		node.Next = MergeSortList(l1.Next, l2)
	}
	return node
}
