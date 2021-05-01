package leetcode

// 层遍历取最右边的数据
func rightSideView(root *TreeNode) []int {
	r := make([]int, 0)
	if root == nil {
		return r
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		right := queue[len(queue)-1]
		r = append(r, right.Val)
		tmp := make([]*TreeNode, 0)
		for _, node := range queue {
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		queue = tmp
	}
	//rightSideViewAux(root, &r)
	return r
}

func rightSideViewAux(root *TreeNode, r *[]int) {
	if root == nil {
		return
	}
	*r = append(*r, root.Val)
	rightSideViewAux(root.Right, r)
}
