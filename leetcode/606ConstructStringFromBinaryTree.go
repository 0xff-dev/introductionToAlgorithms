package leetcode

import (
	"fmt"
	"strings"
)

/*
	 1
  2 .  3
4
1(2(4))(3)
*/
func tree2str(root *TreeNode) string {
	sb := strings.Builder{}
	buildStringFromTree(root, &sb)
	return sb.String()
}

func buildStringFromTree(root *TreeNode, sb *strings.Builder) {
	if root == nil {
		return
	}
	sb.WriteString(fmt.Sprintf("%d", root.Val))
	if root.Left == nil && root.Right == nil {
		return
	}
	sb.WriteByte('(')
	buildStringFromTree(root.Left, sb)
	sb.WriteByte(')')

	if root.Right != nil {
		sb.WriteByte('(')
		buildStringFromTree(root.Right, sb)
		sb.WriteByte(')')
	}
}
