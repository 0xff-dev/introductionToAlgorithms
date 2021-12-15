package leetcode

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	for pre, walker := head, head.Next; walker != nil; {
		var inPre *ListNode
		found := false
		for inWalker := head; inWalker != walker; inWalker = inWalker.Next {
			if inWalker.Val > walker.Val {
				found = true
				break
			}
			inPre = inWalker
		}
		if found {
			pre.Next = walker.Next
			if inPre == nil {
				walker.Next = head
				head = walker
			} else {
				walker.Next = inPre.Next
				inPre.Next = walker
			}
			walker = pre.Next
			continue
		}
		pre, walker = walker, walker.Next
	}

	return head
}

func insertionSortList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	outPre, outWalker := head, head.Next
	for outWalker != nil {
		in, pre := head, head
		for ; in != outWalker && in.Val < outWalker.Val; in = in.Next {
			pre = in
		}
		if in == outWalker {
			outWalker, outPre = outWalker.Next, outWalker
			continue
		}
		outPre.Next, outWalker.Next = outWalker.Next, in
		if in == head {
			head = outWalker
		} else {
			pre.Next = outWalker
		}
		outWalker = outPre.Next
	}
	return head
}