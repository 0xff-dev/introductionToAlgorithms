package leetcode

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	ans := root.Val
	ansLevel := 1
	level := 1
	for len(queue) > 0 {
		next := make([]*TreeNode, 0)
		tmpSum := 0
		for _, item := range queue {
			if item.Left != nil {
				tmpSum += item.Left.Val
				next = append(next, item.Left)
			}
			if item.Right != nil {
				tmpSum += item.Right.Val
				next = append(next, item.Right)
			}
		}
		if len(next) > 0 && tmpSum > ans {
			ansLevel = level + 1
			ans = tmpSum
		}
		queue = next
		level++
	}
	return ansLevel
}
