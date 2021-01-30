package leetcode

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	l1Array, l2Array := make([]int, 0), make([]int, 0)
	getListIntArray(l1, &l1Array)
	getListIntArray(l2, &l2Array)
	var head *ListNode
	l1Idx, l2Idx := 0, 0
	cf := 0
	for ; l1Idx < len(l1Array) && l2Idx < len(l2Array); l1Idx, l2Idx = l1Idx+1, l2Idx+1 {
		tmpVal := l1Array[l1Idx] + l2Array[l2Idx] + cf
		cf = tmpVal / 10
		if tmpVal > 9 {
			tmpVal -= 10
		}
		newHeadNode := &ListNode{
			Val:  tmpVal,
			Next: head,
		}
		head = newHeadNode
	}
	for ; l1Idx < len(l1Array); l1Idx++ {
		tmpVal := l1Array[l1Idx] + cf
		cf = tmpVal / 10
		if tmpVal > 9 {
			tmpVal -= 10
		}
		newHead := &ListNode{
			Val:  tmpVal,
			Next: head,
		}
		head = newHead
	}

	for ; l2Idx < len(l2Array); l2Idx++ {
		tmpVal := l2Array[l2Idx] + cf
		cf = tmpVal / 10
		if tmpVal > 9 {
			tmpVal -= 10
		}
		newHead := &ListNode{
			Val:  tmpVal,
			Next: head,
		}
		head = newHead
	}
	if cf > 0 {
		newHead := &ListNode{
			Val:  cf,
			Next: head,
		}
		head = newHead
	}
	return head
}

func getListIntArray(head *ListNode, array *[]int) {
	if head == nil {
		return
	}
	getListIntArray(head.Next, array)
	*array = append(*array, head.Val)
}
