package leetcode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := 0, 0
	for walker := headA; walker != nil; walker = walker.Next {
		a++
	}

	for walker := headB; walker != nil; walker = walker.Next {
		b++
	}
	if a > b {
		for steps := 0; steps < a-b; steps++ {
			headA = headA.Next
		}
	} else {
		for steps := 0; steps < b-a; steps++ {
			headB = headB.Next
		}
	}

	for wa, wb := headA, headB; wa != nil && wb != nil; wa, wb = wa.Next, wb.Next {
		if wa == wb {
			return wa
		}
	}
	return nil
}
