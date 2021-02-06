package leetcode

func levelOrder(root *TreeNode) [][]int {
	r := make([][]int, 0)
	if root == nil {
		return r
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		tmpQ := make([]*TreeNode, 0)
		tmpV := make([]int, 0)
		for _, node := range queue {
			tmpV = append(tmpV, node.Val)
			if node.Left != nil {
				tmpQ = append(tmpQ, node.Left)
			}
			if node.Right != nil {
				tmpQ = append(tmpQ, node.Right)
			}
		}

		queue = tmpQ
		r = append(r, tmpV)
	}
	return r
}
