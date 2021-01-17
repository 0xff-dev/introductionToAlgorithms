package leetcode

func rotateRight(head *ListNode, k int) *ListNode {
	length := 0
	for walker := head; walker != nil; walker = walker.Next {
		length++
	}

	if length == 0 || length == 1 {
		return head
	}

	k = k % length
	if k == 0 {
		return head
	}

	walker := head
	for idx := length - k - 1; idx > 0; idx-- {
		walker = walker.Next
	}

	newHead := walker.Next
	walker.Next = nil
	walker = newHead
	for walker.Next != nil {
		walker = walker.Next
	}
	walker.Next = head

	return newHead
}
