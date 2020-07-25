package offer

// 有两种做法，独立思考完成
func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p, walker := head, head.Next
	for walker != nil {
		next := walker.Next
		walker.Next = p
		p = walker
		walker = next
	}
	head.Next = nil
	return p
}

func ReverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node := ReverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}
