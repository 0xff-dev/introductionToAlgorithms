package leetcode

func mergeNodes(head *ListNode) *ListNode {
	// 0, 1, 2, 3, 0,4,5,6,0
	skipNode := head
	walker := head.Next
	for ; walker.Next != nil; walker = walker.Next {
		if walker.Val == 0 {
			skipNode.Next = walker
			skipNode = walker
			continue
		}
		skipNode.Val += walker.Val
	}
	skipNode.Next = nil
	return head
}
