package leetcode

func levelOrder1(root *MultiChildNode) [][]int {
	r := make([][]int, 0)
	if root == nil {
		return r
	}
	level := 0
	queue := []*MultiChildNode{root}
	for len(queue) > 0 {
		r = append(r, []int{})
		nextQueue := make([]*MultiChildNode, 0)
		for _, node := range queue {
			r[level] = append(r[level], node.Val)
			if len(node.Children) > 0 {
				nextQueue = append(nextQueue, node.Children...)
			}
		}
		level++
		queue = nextQueue
	}

	return r
}
