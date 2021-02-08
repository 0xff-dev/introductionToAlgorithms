package leetcode

func generateTrees(n int) []*TreeNode {
	r := make([][]*TreeNode, 9)
	r[0] = []*TreeNode{nil}
	r[1] = []*TreeNode{&TreeNode{1, nil, nil}}
	r[2] = []*TreeNode{&TreeNode{1, nil, &TreeNode{2, nil, nil}}, &TreeNode{2, &TreeNode{1, nil, nil}, nil}}

	for idx := 3; idx <= n; idx++ {
		r[idx] = make([]*TreeNode, 0)

		for topNum := 1; topNum <= idx; topNum++ {
			for _, lChild := range r[topNum-1] {
				for _, rChild := range r[idx-topNum] {
					treeRoot := &TreeNode{topNum, nil, nil}
					treeRoot.Left = lChild
					if rChild != nil {
						tmpRChild := &TreeNode{rChild.Val + topNum, nil, nil}
						copyTree(rChild, tmpRChild, topNum)
						treeRoot.Right = tmpRChild
					}
					r[idx] = append(r[idx], treeRoot)
				}
			}
		}
	}
	return r[n]
}

func copyTree(tree *TreeNode, other *TreeNode, topNum int) {
	if tree == nil {
		return
	}

	if tree.Left != nil {
		other.Left = &TreeNode{tree.Left.Val + topNum, nil, nil}
		copyTree(tree.Left, other.Left, topNum)
	}
	if tree.Right != nil {
		other.Right = &TreeNode{tree.Right.Val + topNum, nil, nil}
		copyTree(tree.Right, other.Right, topNum)
	}

}
