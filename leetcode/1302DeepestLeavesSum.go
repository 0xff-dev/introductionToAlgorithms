package leetcode

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	sum := 0
	for len(queue) > 0 {
		tmp := make([]*TreeNode, 0)
		sum = 0
		for _, node := range queue {
			sum += node.Val
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		queue = tmp
	}
	return sum
}

func solution2(root *TreeNode) int {
	maxDep, sum := 0, 0
	s2aux(root, &maxDep, &sum, 0)
	return sum
}

func s2aux(root *TreeNode, maxDep, sum *int, nowDep int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		// leaves
		if nowDep > *maxDep {
			*maxDep = nowDep
			*sum = root.Val
			return
		} else if nowDep == *maxDep {
			*sum += root.Val
		}
		return
	}
	if root.Left != nil {
		s2aux(root.Left, maxDep, sum, nowDep+1)
	}
	if root.Right != nil {
		s2aux(root.Right, maxDep, sum, nowDep+1)
	}
}
