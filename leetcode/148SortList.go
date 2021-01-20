package leetcode

func sortList(head *ListNode) *ListNode {
	auxSort(head, nil)
	return head
}

func auxSort(start, end *ListNode) {
	if start == end || start.Next == end {
		return
	}
	sign, walker := start, start
	aim := sign.Val
	for walker != end {
		if walker.Val < aim {
			sign = sign.Next
			sign.Val, walker.Val = walker.Val, sign.Val
		}
		walker = walker.Next
	}
	sign.Val, start.Val = start.Val, sign.Val
	auxSort(start, sign)
	auxSort(sign.Next, end)
}
