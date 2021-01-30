package leetcode

import "testing"

func TestMyLinkedList(t *testing.T) {
	//list := Constructor()
	//list.AddAtHead(1)
	//list.AddAtTail(3)
	//t.Log(list.Data)
	//
	//list.AddAtIndex(1, 2)
	//t.Log(list.Data)
	//
	//t.Log(list.Get(1))
	//list.DeleteAtIndex(1)
	//t.Log(list.Get(1))

	l1 := Constructor()
	l1.AddAtHead(2)
	l1.DeleteAtIndex(1)
	l1.AddAtHead(2)
	l1.AddAtHead(7)
	l1.AddAtHead(3)
	l1.AddAtHead(2)
	l1.AddAtHead(5)
	l1.AddAtTail(5)
	t.Log("get 5: ", l1.Get(5))
	t.Log(l1.Data)
	l1.DeleteAtIndex(6)
	l1.DeleteAtIndex(4)

	l2 := Constructor()
	l2.AddAtIndex(0, 2)
	l2.AddAtIndex(1, 3)
	l2.AddAtIndex(2, 5)
	l2.AddAtIndex(3, 8)
	l2.AddAtIndex(4, 10)
	l2.AddAtIndex(5, 12)
	l2.DeleteAtIndex(5)
	t.Log(l2.Data)

	l3 := Constructor()
	l3.AddAtHead(4)
	t.Log("get 1", l3.Get(1))
	l3.AddAtHead(1)
	l3.AddAtHead(5)
	l3.DeleteAtIndex(3)
	t.Log(l3.Data)
}
