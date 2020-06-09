// 二叉搜索树
package datastructure

import "fmt"

func (tNode *treeNode) Search(key int) *treeNode {
	if tNode == nil {
		return nil
	}
	if tNode.Value() == key {
		return tNode
	} else if tNode.Value() < key {
		return tNode.Right.Search(key)
	} else {
		return tNode.Left.Search(key)
	}
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
func GetNode(tNode *treeNode, left bool) (*treeNode, *treeNode) {
	var father *treeNode
	if left {
		for tNode.Right != nil {
			father = tNode
			tNode = tNode.Right
		}
	} else {
		for tNode.Left != nil {
			father = tNode
			tNode = tNode.Left
		}
	}
	return father, tNode
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
		_, node := GetNode(selfNode.Left, true)
		return node
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
		_, node := GetNode(selfNode.Right, false)
		return node
	}
	if rootIndex == nil || firstParent == nil {
		return nil
	}
	if rootIndex.Left == selfNode {
		return rootIndex
	}
	return firstParent
}

func (tNode *treeNode) BinarySearchTree2DoublyList() {
	var head *treeNode
	aux(tNode, &head)
	for walker := &head; *walker != nil; *walker = (*walker).Left {
		fmt.Println("now item: ", (*walker).Data)
	}
}

func aux(root *treeNode, last **treeNode) {
	if root == nil {
		return
	}
	cur := root
	if cur.Left != nil {
		aux(cur.Left, last)
	}
	cur.Left = *last
	if *last != nil {
		(*last).Right = cur
	}
	*last = cur
	if cur.Right != nil {
		aux(cur.Right, last)
	}
}

// 先不管平衡的问题
func (tNode *treeNode) InsertNode(val int) {
	insert(tNode, val)
}

func insert(root *treeNode, val int) {
	walker := root
	var father *treeNode
	var cmp int
	for walker != nil {
		father = walker
		cmp, _ = walker.Data.(int)
		if cmp > val {
			walker = walker.Left
		} else {
			walker = walker.Right
		}
	}
	if father == nil {
		root = &treeNode{Data: val}
		return
	}
	node := &treeNode{Data: val}
	if cmp > val {
		father.Right = node
	} else {
		father.Left = node
	}
}

// 搜索二叉树节点分几种情况
// 1. 没有任何节点的时候直接删除就好
// 2. 只有一个孩子，那么，这个孩子替代他即可
// 3. 有连个孩子，那么需要找目标节点的后继，
// 目前首先有点啰嗦，周末精简
func DeleteNode(root **treeNode, val int) {
	selfNode, father, _ := (*root).FindParent(val, false)
	if selfNode == nil {
		return
	}
	if selfNode.Left == nil {
		if father == nil {
			*root = (*root).Right
			fmt.Println("root", root)
			return
		}
		if selfNode == father.Right {
			father.Right = selfNode.Right
		} else {
			father.Left = selfNode.Right
		}
		return
	}
	if selfNode.Right == nil {
		if father == nil {
			*root = (*root).Left
			return
		}
		if selfNode == father.Right {
			father.Right = selfNode.Left
		} else {
			father.Left = selfNode.Left
		}
		return
	}

	successorFather, successorNode := GetNode(selfNode.Right, false)
	if successorNode != selfNode.Right {
		successorFather.Left = successorNode.Right
		successorNode.Right = selfNode.Right
	}
	successorNode.Left = selfNode.Left
	if father != nil {
		if father.Left == selfNode {
			father.Left = successorNode
		} else {
			father.Right = successorNode
		}
	} else {
		*root = successorNode
	}
}

