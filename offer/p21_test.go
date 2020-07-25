package offer

import (
	"fmt"
	"testing"
)

func TestMinStack(t *testing.T) {
	stack := MinStack{
		DataStack:   &Stack{},
		MinNumStack: &Stack{},
	}

	for _, num := range []int{2, 3, 3, 5, 1, 6} {
		stack.Push(num)
	}
	if min, err := stack.Min(); err == nil && min != 1 {
		t.Fatalf("min number is 1")
	}
	if data, err := stack.Pop(); err == nil {
		fmt.Println("top data 1", data)
	}
	if min, err := stack.Min(); err == nil && min != 1 {
		t.Fatalf("min number is 1")
	}
}
