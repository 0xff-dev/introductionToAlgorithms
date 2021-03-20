package leetcode

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	prePre, pre, walker := head, head, head
	nodeCount := 0
	for ; walker != nil; walker = walker.Next {
		if walker.Val == pre.Val {
			nodeCount++
			continue
		}
		if nodeCount > 1 {
			prePre.Next, nodeCount = walker, 1
			if pre == head {
				head = walker
				prePre = head
			}
			pre = walker
			continue
		}
		prePre, pre = pre, walker
		nodeCount = 1
	}
	if nodeCount > 1 {
		if pre == head {
			return nil
		}
		prePre.Next = nil
	}
	return head
}
