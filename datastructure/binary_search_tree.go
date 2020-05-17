// 二叉搜索树
package datastructure

func (tNode *treeNode) Search(key int) *treeNode {
	compareKey, _ := tNode.Data.(int)
	if tNode.Data == nil || compareKey == key {
		return tNode
	}
	if compareKey < key {
		return tNode.Left.Search(key)
	}
	return tNode.Right.Search(key)
}

// 最小公共祖先
func (tNode *treeNode) LCA(node1, node2 *treeNode) *treeNode {
	if tNode == nil {
		return nil
	}
	if tNode == node1 || tNode == node2 {
		return tNode
	}
	left := tNode.Left.LCA(node1, node2)
	right := tNode.Right.LCA(node1, node2)
	if left != nil && right != nil {
		return tNode
	}
	if left == nil {
		return right
	}
	return left
}

// find predecessor, successor, if tree has parent node
func (tNode *treeNode) FindParentWithParent(val int, predecessor bool) *treeNode {
	if predecessor {
		parent := tNode.Parent
		for parent != nil && parent.Right != tNode {
			tNode = parent
			parent = tNode.Parent
		}
		return parent
	}
	parent := tNode.Parent
	for parent != nil && parent.Left != tNode {
		tNode = parent
		parent = tNode.Parent
	}
	return parent
}

// tree noe does't has parent node
func (tNode *treeNode) FindParent(val int, predecessor bool) (selfNode, rootIndex, firstParent *treeNode) {
	root := tNode
	if predecessor {
		for root != nil {
			intVal, _ := root.Data.(int)
			if intVal == val {
				selfNode = root
				return
			}
			rootIndex = root
			if intVal > val {
				root = root.Left
			} else {
				firstParent = root // 就是满足第三个条件的那个情况 但节点是父亲节点的左子树的时候，需要找到父亲的第一个右转折点
				root = root.Right
			}
		}
	} else {
		for root != nil {
			intVal, _ := root.Data.(int)
			if intVal == val {
				selfNode = root
				return
			}
			rootIndex = root
			if intVal < val {
				root = root.Right
			} else {
				firstParent = root
				root = root.Left
			}
		}
	}
	return
}
func GetNode(tNode *treeNode, left bool) *treeNode {
	if left {
		for tNode.Right != nil {
			tNode = tNode.Right
		}
	} else {
		for tNode.Left != nil {
			tNode = tNode.Left
		}
	}
	return tNode
}

//若一个节点有左子树，那么该节点的前驱节点是其左子树中val值最大的节点（也就是左子树中所谓的rightMostNode）
//若一个节点没有左子树，那么判断该节点和其父节点的关系
//2.1 若该节点是其父节点的右边孩子，那么该节点的前驱结点即为其父节点。
//2.2 若该节点是其父节点的左边孩子，那么需要沿着其父亲节点一直向树的顶端寻找，直到找到一个节点P，P节点是其父节点Q的右边孩子（可参考例子2的前驱结点是1），那么Q就是该节点的后继节点
func (tNode *treeNode) Predecessor(val int) *treeNode {
	selfNode, rootIndex, firstParent := tNode.FindParent(val, true)
	if selfNode == nil {
		return nil // can not find a node by val
	}
	if selfNode.Left != nil {
		return GetNode(selfNode.Left, true)
	}
	if rootIndex == nil || firstParent == nil {
		return nil
	}
	if rootIndex.Right == selfNode {
		return rootIndex
	}
	return firstParent
}

// 如果目标节点的右子树不为空，可以遍历右子树的最大左子树即可
// 如果右子树为空
// 是父节点的左子树，那么后继就是父节点
// 是父节点的右子树，就是A.left=father 的A
func (tNode *treeNode) Successor(val int) *treeNode {
	selfNode, rootIndex, firstParent := tNode.FindParent(val, false)
	if selfNode == nil {
		return nil // not found node by val
	}
	if selfNode.Right != nil {
		return GetNode(selfNode.Right, false)
	}
	if rootIndex == nil || firstParent == nil {
		return nil
	}
	if rootIndex.Left == selfNode {
		return rootIndex
	}
	return firstParent
}
