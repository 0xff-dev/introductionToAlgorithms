package leetcode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head, walker *ListNode
	for l1 != nil && l2 != nil {
		tmpNode := &ListNode{}
		if l1.Val <= l2.Val {
			tmpNode.Val = l1.Val
			l1 = l1.Next
		} else {
			tmpNode.Val = l2.Val
			l2 = l2.Next
		}
		if walker == nil {
			head = tmpNode
			walker = tmpNode
			continue
		}
		walker.Next = tmpNode
		walker = tmpNode
	}
	if l1 != nil {
		walker.Next = l1
	}
	if l2 != nil {
		walker.Next = l2
	}
	return head
}
