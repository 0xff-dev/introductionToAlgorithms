package leetcode

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	walker := &head
	for *walker != nil {
		entry := *walker
		if entry == nil {
			break
		}
		if entry.Val == val {
			*walker = entry.Next
			continue
		}
		walker = &(*walker).Next
	}
	return head
}
