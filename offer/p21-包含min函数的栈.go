package offer

import "fmt"

type Stack struct {
	Data []int
}

func (s *Stack) Push(val int) {
	s.Data = append(s.Data, val)
}
func (s *Stack) Pop() (int, error) {
	if len(s.Data) == 0 {
		return -1, fmt.Errorf("empty stack")
	}
	topData := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return topData, nil
}
func (s *Stack) Top() (int, error) {
	if len(s.Data) == 0 {
		return -1, fmt.Errorf("empty stack")
	}
	return s.Data[len(s.Data)-1], nil
}

type MinStack struct {
	DataStack   *Stack
	MinNumStack *Stack
}

func (ms *MinStack) Push(val int) {
	ms.DataStack.Push(val)
	minNum, err := ms.MinNumStack.Top()
	if err != nil {
		ms.MinNumStack.Push(val)
		return
	}
	if minNum > val {
		minNum = val
	}
	ms.MinNumStack.Push(minNum)
}

func (ms *MinStack) Pop() (int, error) {
	ms.MinNumStack.Pop()
	return ms.DataStack.Pop()
}

func (ms *MinStack) Min() (int, error) {
	return ms.MinNumStack.Top()
}

