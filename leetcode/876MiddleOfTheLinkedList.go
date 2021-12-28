package leetcode

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	one, two := head, head
	for two != nil && two.Next != nil {
		one = one.Next
		two = two.Next.Next
	}

	return one
}
