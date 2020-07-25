package offer

import (
	"fmt"
	"testing"
)

func TestMergeSortList(t *testing.T) {
	list1 := &ListNode{1, &ListNode{3, &ListNode{5, nil}}}
	list2 := &ListNode{2, &ListNode{4, &ListNode{6, nil}}}
	list := MergeSortList(list1, list2)
	for walker := list; walker != nil; walker = walker.Next {
		fmt.Println("Data: ", walker.Val)
	}

	var l1, l2 *ListNode
	list = MergeSortList(l1, l2)
	for walker := list; walker != nil; walker = walker.Next {
		fmt.Println("Data: ", walker.Val)
	}

	l1, l2 = &ListNode{1, nil}, &ListNode{1, nil}
	list = MergeSortList(l1, l2)
	for walker := list; walker != nil; walker = walker.Next {
		fmt.Println("Data: ", walker.Val)
	}
}
