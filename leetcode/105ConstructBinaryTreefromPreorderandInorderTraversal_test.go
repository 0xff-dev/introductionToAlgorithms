package leetcode

import "testing"

func TestBuildTree1(t *testing.T) {
	preOrder, inOrder := []int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}
	tree := buildTree1(preOrder, inOrder)
	r := levelOrder(tree)
	t.Log(r)

	preOrder, inOrder = []int{1}, []int{1}
	tree = buildTree1(preOrder, inOrder)
	r = levelOrder(tree)
	t.Log(r)
}
