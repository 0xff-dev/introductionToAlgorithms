package algorith

import (
	"testing"
)

func TestReverseKNode(t *testing.T) {
	list := &listNode{
		Val:  1,
		Next: nil,
	}
	walker := list
	for key := 2; key < 10; key++ {
		node := &listNode{key, nil}
		walker.Next = node
		walker = node
	}
	list = ReverseKNode(list, 4)
	list.Dis()
}
