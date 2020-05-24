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
	Data   interface{}
	Left   *treeNode
	Right  *treeNode
	Parent *treeNode // this is only for find parent function, in others ways, this var will not be used.
	Depth int
}

func (tNode *treeNode) Print() {
	fmt.Printf("---tree data: %v---\n", tNode.Data)
}

type AvlTree struct {
	Root *treeNode
}

func (node *treeNode) Value() int {
	if node == nil {
		panic(fmt.Errorf("nil pointer"))
	}
	intVal, _ := node.Data.(int)
	return intVal
}
