package leetcode

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	f, s := head, head.Next
	for s != nil && s.Next != nil {
		f = f.Next
		s = s.Next.Next
		if f == s {
			return true
		}
	}

	return false
}
