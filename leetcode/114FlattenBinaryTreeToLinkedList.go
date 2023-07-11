package leetcode

/*
	1

2       3

1

	2
	    3
*/
func flatten114(root *TreeNode) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, nil
	}
	end := root
	rightStart, rightEnd := flatten114(root.Right)
	leftStart, leftEnd := flatten114(root.Left)
	if leftStart != nil {
		end.Right = leftStart
		end = leftEnd
	}
	if rightStart != nil {
		end.Right = rightStart
		end = rightEnd
	}
	root.Left = nil
	return root, end
}
func flatten(root *TreeNode) {
	flatten114(root)
	return
}
