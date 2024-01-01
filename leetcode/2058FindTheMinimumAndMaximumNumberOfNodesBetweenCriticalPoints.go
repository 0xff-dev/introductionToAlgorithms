package leetcode

func nodesBetweenCriticalPoints(head *ListNode) []int {
	result := []int{-1, -1}
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return []int{-1, -1}
	}
	start := -1
	preStart := 1
	end := -1
	pre, cur := head, head.Next
	index := 1
	// 2, 3, 3, 2
	for ; cur.Next != nil; pre, cur, index = cur, cur.Next, index+1 {
		a := pre.Val
		b := cur.Next.Val
		if !(cur.Val > a && cur.Val > b || cur.Val < a && cur.Val < b) {
			continue
		}
		end = index
		if start == -1 {
			start = index
		} else if result[0] == -1 || index-preStart < result[0] {
			result[0] = index - preStart
		}

		preStart = index
	}
	if end != start {
		result[1] = end - start
	}
	return result
}
