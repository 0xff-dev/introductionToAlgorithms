package offer

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	list := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	list = ReverseList(list)
	for walker := list; walker != nil; walker = walker.Next {
		fmt.Println("data: ", walker.Val)
	}
}

func TestReverseListRecursive(t *testing.T) {
	list := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	list = ReverseListRecursive(list)
	for walker := list; walker != nil; walker = walker.Next {
		fmt.Println("data: ", walker.Val)
	}
}
