package challengeprogrammingdatastructure

import (
	"fmt"
	"testing"
)

func TestBinarySearchTreeDelete(t *testing.T) {
	var tree *BSTNode
	values := []int{30, 88, 12, 1, 20, 17, 25}
	for _, v := range values {
		tree = BinarySearchTreeInsert(tree, v)
	}
	inOrder92(tree)
	fmt.Println()

	for _, v := range values {
		t.Logf("delte node %v\n", v)
		BinarySearchTreeDelete(&tree, v)
		inOrder92(tree)
		fmt.Println()
	}
}
