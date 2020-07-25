package offer

func IsPopOrder(order, in []int) bool {
	s1 := &Stack{}
	for index := len(order) - 1; index >= 0; index-- {
		s1.Push(order[index])
	}
	s2 := &Stack{}
	for index := 0; index < len(in); index++ {
		top, _ := s1.Top()
		s2.Push(in[index])
		top2 := in[index]
		for top == top2 && top != -1 {
			top, _ = s1.Top()
			top2, _ = s2.Top()
			if top == top2 {
				s1.Pop()
				s2.Pop()
			}
		}
	}
	return len(s1.Data) == 0
}
