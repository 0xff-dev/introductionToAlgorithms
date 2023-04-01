package challengeprogrammingdatastructure

func SuccessorWithFather(root *BSTNode) (*BSTNode, *BSTNode) {
	if root == nil {
		return nil, nil
	}
	var parent *BSTNode
	walker := root
	for walker.Left != nil {
		parent = walker
		walker = walker.Left
	}
	return parent, walker
}

// 找到一个节点,父节点,中序遍历的前驱节点，中序遍历的后继节点
func FindSelfAndParent(root *BSTNode, val int) (*BSTNode, *BSTNode, *BSTNode, *BSTNode) {
	walker := root
	var self, parent, p, s *BSTNode
	for walker != nil {
		if walker.Val == val {
			self = walker
			break
		}
		parent = walker
		if walker.Val < val {
			p = walker
			walker = walker.Right
		} else {
			s = walker
			walker = walker.Left
		}
	}
	return self, parent, p, s
}

func BinarySearchTreeDelete(root **BSTNode, target int) {
	self, parent, _, _ := FindSelfAndParent(*root, target)

	if self == nil {
		return
	}
	if self.Left == nil {
		if parent == nil {
			*root = (*root).Right
			return
		}
		if self == parent.Left {
			parent.Left = self.Right
		} else {
			parent.Right = self.Right
		}
		return
	}
	if self.Right == nil {
		if parent == nil {
			*root = (*root).Left
			return
		}
		if self == parent.Left {
			parent.Left = self.Left
		} else {
			parent.Right = self.Left
		}
		return
	}

	successorParent, successorNode := SuccessorWithFather(self.Right)
	// 使用后继节点的逻辑
	// 1. 我们需要用后继节点替换self(要被删除的)节点
	if successorNode != self.Right {
		successorParent.Left = successorNode.Right
		successorNode.Right = self.Right
	}

	successorNode.Left = self.Left

	if parent != nil {
		if self == parent.Left {
			parent.Left = successorNode
		} else {
			parent.Right = successorNode
		}
	} else {
		*root = successorNode
	}
}
