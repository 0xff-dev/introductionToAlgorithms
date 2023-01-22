package leetcode

func closestNodes(root *TreeNode, queries []int) [][]int {
	ans := make([][]int, len(queries))
	nodes := make([]int, 0)
	var inOrder func(*TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrder(root.Left)
		nodes = append(nodes, root.Val)
		inOrder(root.Right)
	}
	inOrder(root)
	for i, q := range queries {
		l, r := binarySearch(nodes, q)
		ans[i] = []int{l, r}
	}

	return ans
}

func findClosestNodes(root *TreeNode, target int, left, right *int) {
	if root == nil {
		return
	}
	if root.Val == target {
		*left, *right = root.Val, root.Val
		return
	}
	if root.Val > target {
		*right = root.Val
		findClosestNodes(root.Left, target, left, right)
		return
	}
	*left = root.Val
	findClosestNodes(root.Right, target, left, right)
}

// 草，二分搜索
func binarySearch(nodes []int, target int) (int, int) {
	l, r := 0, len(nodes)-1

	left, right := -1, -1
	for l <= r {
		mid := (r-l)/2 + l
		if nodes[mid] == target {
			return target, target
		}
		if nodes[mid] < target {
			left = nodes[mid]
			l = mid + 1
			continue
		}
		right = nodes[mid]
		r = mid - 1
	}
	return left, right
}
