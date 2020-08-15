package offer

func getListLength(list *ListNode) int {
	walker := list
	length := 0
	for ; walker != nil; walker = walker.Next {
		length++
	}
	return length
}

func CommonNode(list1, list2 *ListNode) *ListNode {
	if list1 == nil || list2 == nil {
		return nil
	}
	l1, l2 := getListLength(list1), getListLength(list2)
	walker1, walker2 := list1, list2
	if l1 > l2 {
		for step := l1 - l2; step > 0; step-- {
			walker1 = walker1.Next
		}
	} else {
		for step := l2 - l1; step > 0; step-- {
			walker2 = walker2.Next
		}
	}
	for walker1 != nil && walker2 != nil && walker1 != walker2 {
		walker1, walker2 = walker1.Next, walker2.Next
	}
	if walker1 == nil || walker2 == nil {
		return nil
	}
	return walker2
}
