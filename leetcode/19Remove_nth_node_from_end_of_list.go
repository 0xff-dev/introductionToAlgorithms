package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	newHead, walker := head, head
	for idx := 0; walker != nil && idx < n; idx++ {
		walker = walker.Next
	}
	if walker == nil {
		head = head.Next
		return head
	}

	for walker.Next != nil {
		walker, newHead = walker.Next, newHead.Next
	}
	newHead.Next = newHead.Next.Next
	return head
}
