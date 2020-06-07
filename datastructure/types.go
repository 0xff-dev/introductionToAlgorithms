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
	Depth  int
	Color  bool
}

func (tNode *treeNode) Print() {
	fmt.Printf("---tree data: %v--- color: %s\n", tNode.Data, tNode.color())
}

func (tNode *treeNode) color() string{
	if tNode.Color {
		return "RED"
	}
	return "BLACK"
}
type AvlTree struct {
	Root *treeNode
}

type RBTree struct {
	Root *treeNode
}

const (
	RED   = true
	BLACk = false
)

func (node *treeNode) Value() int {
	if node == nil {
		panic(fmt.Errorf("nil pointer"))
	}
	intVal, _ := node.Data.(int)
	return intVal
}
