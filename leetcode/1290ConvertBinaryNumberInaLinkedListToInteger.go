package leetcode

func getDecimalValue(head *ListNode) int {
	if head == nil {
		return 0
	}
	base := 1
	return traversalList(head, &base)
}

func traversalList(head *ListNode, base *int) int {
	if head == nil {
		return 0
	}
	now := traversalList(head.Next, base)
	now = head.Val*(*base) + now
	*base <<= 1
	return now
}
