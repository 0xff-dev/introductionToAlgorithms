package leetcode

func splitListToParts(head *ListNode, k int) []*ListNode {
	l := 0
	for walker := head; walker != nil; walker = walker.Next {
		l++
	}

	ans := make([]*ListNode, k)
	if k >= l {
		walker := head
		index := 0
		for walker != nil {
			next := walker.Next
			walker.Next = nil
			ans[index] = walker
			walker = next
			index++
		}
		return ans
	}

	avg := l / k
	left := l % k
	walker := head
	for idx := 0; idx < k; idx++ {
		selectNode := avg
		if idx < left {
			selectNode++
		}
		ans[idx] = walker
		pre := walker
		for selectNode > 0 {
			pre = walker
			walker = walker.Next
			selectNode--
		}
		pre.Next = nil
	}

	return ans
}
