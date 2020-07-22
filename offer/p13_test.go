package offer

import (
	"fmt"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	list := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	DeleteNode(&list.Next.Next)
	for walker := list; walker != nil; walker = walker.Next {
		fmt.Println("data: ", walker.Val)
	}
}
