package leetcode

import "testing"

func TestBuildTree(t *testing.T) {
	inOrder, postOrder := []int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}
	tree := buildTree(inOrder, postOrder)
	r := levelOrder(tree)
	t.Log(r)

	inOrder, postOrder = []int{1}, []int{1}
	tree = buildTree(inOrder, postOrder)
	r = levelOrder(tree)
	t.Log(r)

}
