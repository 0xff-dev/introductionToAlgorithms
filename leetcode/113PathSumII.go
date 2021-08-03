package leetcode

func pathSum(root *TreeNode, targetSum int) [][]int {
	r := make([][]int, 0)

	pathSumIIAux(root, 0, targetSum, []int{}, &r)
	return r
}

func pathSumIIAux(root *TreeNode, now, targetSum int, path []int, r *[][]int) {
	if root == nil {
		return
	}

	now += root.Val
	if root.Left == nil && root.Right == nil {
		if now == targetSum {
			dst := make([]int, len(path))
			copy(dst, path)
			dst = append(dst, root.Val)
			*r = append(*r, dst)
		}
		return
	}

	pathSumIIAux(root.Left, now, targetSum, append(path, root.Val), r)
	pathSumIIAux(root.Right, now, targetSum, append(path, root.Val), r)
}
