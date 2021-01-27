package leetcode

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	second := head.Next
	w1, w2 := head, second
	flag := 1

	for {
		flag = 1 - flag
		if flag == 0 {
			if w2 == nil {
				break
			}

			next := w2.Next
			w1.Next = next
			if next == nil {
				w1.Next = second
				break
			}
			w1 = next
			continue
		}

		next := w1.Next
		w2.Next = next
		if next == nil {
			w1.Next = second
			break
		}
		w2 = next
	}
	return head
}
