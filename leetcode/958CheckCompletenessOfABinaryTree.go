package leetcode

func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	nowLackNode := false
	for len(queue) > 0 {
		nextQ := make([]*TreeNode, 0)
		lackNode := false
		for _, item := range queue {
			if lackNode {
				if item.Left != nil || item.Right != nil {
					return false
				}
				continue
			}
			if item.Left == nil && item.Right != nil {
				return false
			}
			if item.Left != nil {
				nextQ = append(nextQ, item.Left)
			}
			if item.Right != nil {
				nextQ = append(nextQ, item.Right)
			}
			lackNode = item.Right == nil
		}

		// 当前队列的一层与下一层的判断
		if nowLackNode && len(nextQ) > 0 {
			return false
		}
		nowLackNode = lackNode
		queue = nextQ
	}
	return true
}
