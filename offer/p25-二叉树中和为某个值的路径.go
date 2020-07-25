package offer

import "fmt"

func FindPath(tree *TreeNode, value int) {
	path := make([]int, 0)
	aux(tree, value, 0, path)
}

func aux(tree *TreeNode, value, sum int, path []int) {
	if tree == nil {
		return
	}
	tmp := sum + tree.Val
	if tmp == value && tree.Left == nil && tree.Right == nil {
		fmt.Println("find path", append(path, tree.Val))
		return
	}
	if tree.Left != nil {
		aux(tree.Left, value, tmp, append(path, tree.Val))
	}
	if tree.Right != nil {
		aux(tree.Right, value, tmp, append(path, tree.Val))
	}
}
