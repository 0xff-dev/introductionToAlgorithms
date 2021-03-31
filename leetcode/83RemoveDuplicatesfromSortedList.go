package leetcode

func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	walker := &head
	for *walker != nil {
		entry := *walker
		if entry.Next == nil {
			break
		}
		if entry.Next.Val == entry.Val {
			*walker = entry.Next
			continue
		}
		walker = &(*walker).Next
	}
	return head
}
