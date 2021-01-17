package leetcode

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	first := true
	s, e := head, head.Next
	var pre *ListNode
	for e != nil {

		nextNode := e.Next
		s, e = e, s
		s.Next, e.Next = e, nextNode
		if pre != nil {
			pre.Next = s
		}

		if first {
			// set head
			first = false
			head = s
		}
		pre = e
		s = nextNode
		if s == nil {
			break
		}
		e = s.Next
	}
	return head
}
