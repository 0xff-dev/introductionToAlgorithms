package challengeprogrammingdatastructure

import "testing"

func TestBinarySearchTreeSearch(t *testing.T) {
	var tree *BSTNode
	tree = BinarySearchTreeInsert(tree, 30)
	tree = BinarySearchTreeInsert(tree, 88)
	tree = BinarySearchTreeInsert(tree, 12)
	tree = BinarySearchTreeInsert(tree, 1)
	tree = BinarySearchTreeInsert(tree, 20)
	tree = BinarySearchTreeInsert(tree, 17)
	tree = BinarySearchTreeInsert(tree, 25)

	if r := BinarySearchTreeSearch(tree, 12); !r {
		t.Fatalf("find 12 expect true get false")
	}

	if r := BinarySearchTreeSearch(tree, 16); r {
		t.Fatalf("find 16 expect false get true")
	}
}
