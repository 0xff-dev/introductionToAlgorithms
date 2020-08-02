package offer

import "fmt"

type ComplexListNode struct {
	Val   int
	Next  *ComplexListNode
	Other *ComplexListNode
}

func (cln *ComplexListNode) Print() {
	for walker := cln; walker != nil; walker = walker.Next {
		fmt.Print("list node val: ", walker.Val)
		if walker.Other == nil {
			fmt.Print(" --: other is nil\n")
			continue
		}
		fmt.Printf(" --: other is %d\n", walker.Other.Val)
	}
}

func ComplexListReplication(list *ComplexListNode) *ComplexListNode {
	if list == nil {
		return list
	}
	cloneList(list)
	list.Print()
	connectList(list)
	list.Print()
	return combineList(list)
}

// get double node
func cloneList(list *ComplexListNode) {
	walker := list
	for walker != nil {
		next := walker.Next
		tmpNode := &ComplexListNode{
			Val:   walker.Val,
			Next:  nil,
			Other: nil,
		}
		tmpNode.Next = next
		walker.Next = tmpNode
		walker = next
	}
}

func connectList(list *ComplexListNode) {
	walker := list
	for walker != nil {
		next := walker.Next
		if walker.Other != nil {
			next.Other = walker.Other.Next
		}
		walker = walker.Next.Next
	}
}

func combineList(list *ComplexListNode) *ComplexListNode {
	if list == nil {
		return nil
	}
	sourceList, newList, head := list, list.Next, list.Next
	for sourceList != nil {
		sourceList.Next = newList.Next
		sourceList = sourceList.Next
		if sourceList != nil {
			newList.Next = sourceList.Next
			newList = newList.Next
		}
	}
	return head
}
