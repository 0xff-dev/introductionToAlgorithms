package leetcode

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	loop, index, start := 1, head, head
	var newHead, preEnd *ListNode
	for index != nil {
		if loop == k {
			next := index.Next
			tmpHead, tmpEnd := reverseAux(start, index)
			if newHead == nil {
				newHead = tmpHead
			}
			if preEnd != nil {
				preEnd.Next = tmpHead
			}
			preEnd = tmpEnd
			tmpEnd.Next, index, start = next, next, next
			loop = 1
			continue
		}
		index = index.Next
		loop++
	}
	if newHead == nil {
		return head
	}
	return newHead
}

// reverseAux return start and end node
func reverseAux(start, end *ListNode) (*ListNode, *ListNode) {
	p, q := start, start.Next
	p.Next = nil
	for p != end {
		next := q.Next
		q.Next = p
		p = q
		q = next
	}
	return p, start
}
