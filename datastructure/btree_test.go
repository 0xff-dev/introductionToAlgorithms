package datastructure

import "testing"

func TestBtreeInsert(t *testing.T) {
	tree := bTree{Root: newbTreeNode()}
	testNums := []int{100, 34, 78, 234, 16, 72384, 77, 123, 7878, 643, 1273, 5782, 12, 1, 0, 346, 73842, 669, 400, 25, 189, 132, 3743, 1237, 7989, 12378, 78, 818, 623, 634832, 90, 6474, 1234, 7873}
	for i := 0; i < len(testNums); i++ {
		tree.Insert(testNums[i])
	}
	tree.Root.String()
	n := bTreeSearch(tree.Root, 5782)
	if n == nil {
		t.Fatal("can't find aim 3, but it exists")
	}
}
