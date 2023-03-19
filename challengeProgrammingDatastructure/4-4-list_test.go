package challengeprogrammingdatastructure

import "testing"

func TestDoublyLinkedList(t *testing.T) {
	var list *DLNode
	for i := 0; i < 3; i++ {
		t.Logf("insert %d\n", i)
		list = insert(list, i+1)
		list.Print(false)
		list.Print(true)
	}

	t.Log("insert second 3\n")
	list = insert(list, 3)
	list.Print(false)
	list.Print(true)

	for i := 4; i >= 0; i-- {
		t.Logf("delete first %d\n", i)
		list = delete(list, i)
		list.Print(false)
		list.Print(true)
	}
	t.Logf("delete the last 3")
	list = delete(list, 3)
	list.Print(false)
	list.Print(true)

	t.Log("append 0, 1,2, 3")
	for i := 0; i < 4; i++ {
		list = insert(list, i)
	}
	list.Print(false)
	list.Print(true)
	for i := 3; i >= 0; i-- {
		t.Logf("delete first node %d\n", i)
		list = deleteFirst(list)
		list.Print(false)
		list.Print(true)
	}

	for i := 0; i < 4; i++ {
		list = insert(list, i)
	}
	list.Print(false)
	list.Print(true)

	for i := 0; i < 4; i++ {
		t.Logf("delete last node %d\n", i)
		list = deleteLast(list)
		list.Print(false)
		list.Print(true)
	}
}
