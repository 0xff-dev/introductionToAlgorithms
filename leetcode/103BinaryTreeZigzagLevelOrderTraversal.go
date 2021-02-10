package leetcode

func zigzagLevelOrder(root *TreeNode) [][]int {
	r := make([][]int, 0)
	if root == nil {
		return r
	}

	flag := 1
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		tmpQueue := make([]*TreeNode, 0)
		tmpArr := make([]int, 0)
		flag = 1 - flag
		if flag == 0 {
			// left to right
			for idx := 0; idx < len(queue); idx++ {
				tmpArr = append(tmpArr, queue[idx].Val)
				if queue[idx].Left != nil {
					tmpQueue = append(tmpQueue, queue[idx].Left)
				}

				if queue[idx].Right != nil {
					tmpQueue = append(tmpQueue, queue[idx].Right)
				}
			}
			r = append(r, tmpArr)
			queue = tmpQueue
			continue
		}
		lenQueue := len(queue)
		for idx := lenQueue - 1; idx >= 0; idx-- {
			tmpArr = append(tmpArr, queue[idx].Val)
			if queue[lenQueue-idx-1].Left != nil {
				tmpQueue = append(tmpQueue, queue[lenQueue-idx-1].Left)
			}
			if queue[lenQueue-idx-1].Right != nil {
				tmpQueue = append(tmpQueue, queue[lenQueue-idx-1].Right)
			}
		}

		r = append(r, tmpArr)
		queue = tmpQueue
	}
	return r
}
