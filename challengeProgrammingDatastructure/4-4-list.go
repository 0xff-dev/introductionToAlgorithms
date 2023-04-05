package challengeprogrammingdatastructure

import (
	"fmt"
)

func (l *DLNode) Print(reverse bool) {
	if l == nil {
		fmt.Println("list is nil.")
		return
	}
	walker := l
	if reverse {
		for {
			fmt.Printf("%d ", walker.val)
			if walker.pre == l {
				break
			}
			walker = walker.pre
		}
	} else {
		for {
			fmt.Printf("%d ", walker.val)
			if walker.next == l {
				break
			}
			walker = walker.next
		}
	}
	fmt.Println()

}

func insert(list *DLNode, x int) *DLNode {
	node := &DLNode{val: x}
	if list == nil {
		node.next = node
		node.pre = node
	} else {
		pre := list.pre
		pre.next = node
		node.pre = pre
		pre.next = node
		list.pre = node
		node.next = list
	}
	return node
}

func delete(list *DLNode, x int) *DLNode {
	if list.next == list {
		if list.val == x {
			return nil
		}
		return list
	}
	walker := list
	for {
		if walker.val == x {
			walker.pre.next = walker.next
			walker.next.pre = walker.pre
			if walker == list {
				return list.next
			}
			return list
		}
		if walker.next == list {
			return list
		}
		walker = walker.next
	}
}

func deleteFirst(list *DLNode) *DLNode {
	if list == nil {
		return list
	}
	if list.next == list {
		return nil
	}
	list.pre.next = list.next
	list.next.pre = list.pre
	return list.next
}

func deleteLast(list *DLNode) *DLNode {
	if list == nil {
		return list
	}
	if list.next == list {
		return nil
	}
	lastNode := list.pre
	lastNode.pre.next = lastNode.next
	lastNode.next.pre = lastNode.pre
	return list
}
