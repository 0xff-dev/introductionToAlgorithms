package datastructure

import (
	"go-common/library/log"
)

func NewRBTree() *RBTree {
	return &RBTree{
		Root: nil,
	}
}

func (rb *RBTree) rbLeftRotate(node *treeNode) {
	right := node.Right
	node.Right = right.Left
	if right.Left != nil {
		right.Left.Parent = node
	}
	right.Parent = node.Parent
	if node.Parent == nil {
		rb.Root = right
	} else {
		if node == node.Parent.Left {
			node.Parent.Left = right
		} else {
			node.Parent.Right = right
		}
	}
	right.Left = node
	node.Parent = right
}

func (rb *RBTree) rbRightRotate(node *treeNode) {
	left := node.Left
	node.Left = left.Right
	if left.Right != nil {
		left.Right.Parent = node
	}
	left.Parent = node.Parent
	if node.Parent == nil {
		rb.Root = left
	} else {
		if node == node.Parent.Left {
			node.Parent.Left = left
		} else {
			node.Parent.Right = left
		}
	}
	left.Right = node
	node.Parent = left
}

func (rb *RBTree) InsertFixup(node *treeNode) {
	for node.Parent != nil && node.Parent.Color == RED {
		grandFather := node.Parent.Parent
		if node.Parent == grandFather.Left {
			uncle := grandFather.Right
			if uncle != nil && uncle.Color == RED {
				grandFather.Color = RED
				uncle.Color = BLACk
				node.Parent.Color = BLACk
				node = grandFather
			} else {
				if node == node.Parent.Right {
					node = node.Parent
					rb.rbLeftRotate(node)
					grandFather = node.Parent.Parent
				}
				node.Parent.Color = BLACk
				grandFather.Color = RED
				rb.rbRightRotate(grandFather)
			}
		} else {
			// mirror, right children
			uncle := grandFather.Left
			if uncle != nil && uncle.Color == RED {
				grandFather.Color = RED
				uncle.Color = BLACk
				node.Parent.Color = BLACk
				node = grandFather
			} else {
				if node == node.Parent.Left {
					node = node.Parent
					rb.rbRightRotate(node)
					grandFather = node.Parent.Parent
				}
				node.Parent.Color = BLACk
				grandFather.Color = RED
				rb.rbLeftRotate(grandFather)
			}
		}
	}
	rb.Root.Color = BLACk
}

func (rb *RBTree) Insert(val int) {
	node := &treeNode{
		Data:   val,
		Color:  RED,
	}
	cur := rb.Root
	var father *treeNode
	for cur != nil {
		father = cur
		if val < cur.Value() {
			cur = cur.Left
		} else  if val > cur.Value() {
			cur = cur.Right
		} else {
			log.Info("duplicate val")
			return
		}
	}
	if father == nil {
		rb.Root = node
	} else {
		if val < father.Value() {
			father.Left = node
		} else {
			father.Right = node
		}
	}
	node.Parent = father
	rb.InsertFixup(node)
}

// root
func (rb *RBTree) Delete(val int) {

}
