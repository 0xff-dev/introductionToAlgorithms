package leetcode

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	var lessList, greaterList, newHead, greaterHead *ListNode
	for walker := head; walker != nil; walker = walker.Next {
		if walker.Val < x {
			if lessList == nil {
				lessList = walker
				newHead = lessList
				continue
			}
			lessList.Next = walker
			lessList = lessList.Next
			continue
		}
		if greaterList == nil {
			greaterList = walker
			greaterHead = greaterList
			continue
		}
		greaterList.Next = walker
		greaterList = greaterList.Next
	}
	if greaterList != nil {
		greaterList.Next = nil
	}
	if lessList == nil {
		return greaterHead
	}
	lessList.Next = greaterHead
	return newHead
}
