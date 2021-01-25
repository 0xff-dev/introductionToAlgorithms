package leetcode

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	hasCycle := false
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			hasCycle = true
			break
		}
	}

	if hasCycle {
		fast = head
		for ; fast != slow; fast, slow = fast.Next, slow.Next {
		}
		return fast
	}

	return nil
}
