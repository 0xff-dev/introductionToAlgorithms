package datastructure

import "fmt"

type Compare interface {
	Less(other interface{}) bool
	Swap(other interface{})
	Print()
}

type listNode struct {
	data Compare
	next *listNode
}

type List struct {
	Size int // list length
	Head *listNode
}

func (node *listNode) reverse() *listNode {
	if node == nil || node.next == nil {
		return node
	}
	newNode := node.next.reverse()
	node.next.next = node
	node.next = nil
	return newNode
}

// tree part
type treeNode struct {
	Data  interface{}
	Left  *treeNode
	Right *treeNode
}

func (tNode *treeNode) Print() {
	fmt.Printf("---tree data: %v---\n", tNode.Data)
}
