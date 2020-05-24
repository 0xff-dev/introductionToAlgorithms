package datastructure

func (node *treeNode) heigh() int {
	if node == nil {
		return 0
	}
	return node.Depth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (node *treeNode) rightRotate() *treeNode {
	left := node.Left
	leftRight := left.Right
	left.Right = node
	node.Left = leftRight
	node.Depth = max(node.Left.heigh(), node.Right.heigh()) + 1
	left.Depth = max(left.Left.heigh(), left.Right.heigh()) + 1
	return left
}

func (node *treeNode) leftRotate() *treeNode {
	right := node.Right
	rightLeft := right.Left
	right.Left = node
	node.Right = rightLeft
	node.Depth = max(node.Left.heigh(), node.Right.heigh()) + 1
	right.Depth = max(right.Left.heigh(), right.Right.heigh()) + 1
	return right
}

// 插入节点是不平衡节点的左子树的左边，那么直接右旋
// 插入节点是不平衡节点的左子树的右边，先左旋左子树，在右旋不平衡节点
// 插入节点是不平衡节点的右子树的右边，直接左旋
// 插入节点是不平衡节点的右子树的左边，先右旋，在左旋不平衡节点
func (node *treeNode) insert(val int) *treeNode {
	newNode := &treeNode{
		Data:  val,
		Left:  nil,
		Right: nil,
		Depth: 1,
	}
	if node == nil {
		return newNode
	}
	if node.Value() < val {
		node.Right = node.Right.insert(val)
	} else if val < node.Value() {
		node.Left = node.Left.insert(val)
	} else {
		return node
	}
	node.Depth = 1 + max(node.Left.heigh(), node.Right.heigh())
	balance := node.balance()
	if balance > 1 && val < node.Left.Value() {
		return node.rightRotate()
	}
	if balance > 1 && val > node.Left.Value() {
		node.Left = node.Left.leftRotate()
		return node.rightRotate()
	}
	if balance < -1 && val > node.Right.Value() {
		return node.leftRotate()
	}
	if balance < -1 && val < node.Right.Value() {
		node.Right = node.Right.rightRotate()
		return node.leftRotate()
	}
	return node
}

func (node *treeNode) delete(val int) *treeNode {
	if node == nil {
		return node
	}
	if val < node.Value() {
		node.Left = node.Left.delete(val)
	} else if val > node.Value() {
		node.Right = node.Right.delete(val)
	} else {
		if node.Left != nil && node.Right != nil {
			_, tmpNode := GetNode(node.Right, false)
			node.Data = tmpNode.Data
			node.Right = node.Right.delete(tmpNode.Value())
		} else if node.Left != nil {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	if node == nil {
		return nil
	}
	node.Depth = max(node.Left.heigh(), node.Right.heigh()) + 1
	balance := node.balance()
	if balance > 1 && node.Left.balance() >= 0 {
		return node.rightRotate()
	}
	if balance > 1 && node.Left.balance() <= 0 {
		node.Left = node.Left.leftRotate()
		return node.rightRotate()
	}
	if balance < -1 && node.Right.balance() <= 0 {
		return node.leftRotate()
	}
	if balance < -1 && node.Right.balance() >= 0 {
		node.Right = node.Right.rightRotate()
		return node.leftRotate()
	}
	return node
}

func (node *treeNode) balance() int {
	if node == nil {
		return 0
	}
	return node.Left.heigh() - node.Right.heigh()
}

func NewAvlTree() *AvlTree {
	return &AvlTree{Root: nil}
}

func (tree *AvlTree) Insert(val int) {
	tree.Root = tree.Root.insert(val)
}

func (tree *AvlTree) DeleteNode(val int) {
	tree.Root = tree.Root.delete(val)
}

func (tree *AvlTree) Display() {
	tree.Root.Floor()
}
