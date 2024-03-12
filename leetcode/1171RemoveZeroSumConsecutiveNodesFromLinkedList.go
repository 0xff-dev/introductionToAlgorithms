package leetcode

func removeZeroSumSublists(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	w := head
	cur := 0
	cache := map[int]*ListNode{}
	for w != nil {
		cur += w.Val
		if cur == 0 {
			head = w.Next
			w = head
			cache = map[int]*ListNode{}
			continue
		}
		if v, ok := cache[cur]; ok {
			tmp := cur
			for n := v.Next; n != w; n = n.Next {
				tmp += n.Val
				delete(cache, tmp)
			}
			v.Next = w.Next
			w = v.Next
			continue
		}
		cache[cur] = w
		w = w.Next
	}
	return head
}
