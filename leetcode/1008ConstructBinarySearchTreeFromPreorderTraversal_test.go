package leetcode

import "testing"

func TestBstFromPreorder(t *testing.T) {
	preOrder := []int{8, 5, 1, 7, 10, 12}
	tree := bstFromPreorder(preOrder)
	tree.Floor()

	preOrder = []int{1}
	tree = bstFromPreorder(preOrder)
	tree.Floor()
}
