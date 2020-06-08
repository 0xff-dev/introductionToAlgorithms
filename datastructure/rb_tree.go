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

func (rb *RBTree) DeleteFixup(node *treeNode) {
	for (node == nil || node.Color == BLACk) && node != rb.Root {
		if node == node.Parent.Left {
			// left children
			brother := node.Parent.Right
			if brother.Color == RED {
				brother.Color = BLACk
				rb.rbLeftRotate(node.Parent)
				node.Parent.Color = RED
				brother = node.Parent.Right
			}
			if (brother.Left == nil || brother.Left.Color == BLACk) && (brother.Right == nil || brother.Right.Color == BLACk) {
				brother.Color = RED
				node = node.Parent
			} else {
				if brother.Right ==nil || brother.Right.Color == BLACk {
					if brother.Left != nil {
						brother.Left.Color = BLACk
					}
					brother.Color = RED
					rb.rbRightRotate(brother)
					brother = node.Parent.Right
				}
				brother.Color = node.Parent.Color
				node.Parent.Color = BLACk
				brother.Right.Color = BLACk
				rb.rbLeftRotate(node.Parent)
				node = rb.Root
			}
		} else {
			// right children
			brother := node.Parent.Left
			if brother.Color == RED {
				brother.Color = BLACk
				node.Parent.Color = RED
				rb.rbRightRotate(node.Parent)
				brother = node.Left
			}
			if (brother.Left == nil || brother.Left.Color == BLACk) && (brother.Right == nil || brother.Right.Color == BLACk) {
				brother.Color = RED
				node = node.Parent
			} else {
				if brother.Left == nil || brother.Left.Color == BLACk {
					if brother.Right != nil {
						brother.Right.Color =BLACk
					}
					brother.Color = RED
					rb.rbLeftRotate(brother)
					brother = node.Parent
				}
				brother.Color = node.Parent.Color
				brother.Left.Color = BLACk
				node.Parent.Color = BLACk
				rb.rbRightRotate(node.Parent)
				node = rb.Root
			}
		}
	}
	if node != nil {
		node.Color = BLACk
	}
}

// root
func (rb *RBTree) Delete(node *treeNode) {
	if node == nil {
		return
	}
	if node.Left == nil || node.Right == nil {

 	}
 	// find successor, mark node color
 	// 拼接好father和successor，y的关系

 	// successor.coloe == black
}
