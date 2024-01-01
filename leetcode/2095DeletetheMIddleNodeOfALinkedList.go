package leetcode

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	pre, cur := head, head
	for cur != nil && cur.Next != nil {
		pre = pre.Next
		cur = cur.Next.Next
	}
	if pre.Next == nil {
		head.Next = nil
	} else {
		*pre = *(pre.Next)
	}
	return head
}
