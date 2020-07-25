package offer

func MirrorBinaryTree(bt *TreeNode) {
	if bt == nil {
		return
	}
	bt.Right, bt.Left = bt.Left, bt.Right
	if bt.Right != nil {
		MirrorBinaryTree(bt.Right)
	}
	if bt.Left != nil {
		MirrorBinaryTree(bt.Left)
	}
}
