package challengeprogrammingdatastructure

import "testing"

func TestBinarySearchTreeInsert(t *testing.T) {
	var tree *BSTNode
	tree = BinarySearchTreeInsert(tree, 30)
	tree = BinarySearchTreeInsert(tree, 88)
	tree = BinarySearchTreeInsert(tree, 12)
	tree = BinarySearchTreeInsert(tree, 1)
	tree = BinarySearchTreeInsert(tree, 20)
	tree = BinarySearchTreeInsert(tree, 17)
	tree = BinarySearchTreeInsert(tree, 25)
	inOrder92(tree)
	t.Log("\n")
	preOrder92(tree)
}
