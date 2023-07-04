package leetcode

import "testing"

func TestRecoverFromPreorder(t *testing.T) {
	traversal := "1-2--3--4-5--6--7"
	tree := recoverFromPreorder(traversal)
	tree.Floor()
	traversal = "1-2--3---4-5--6---7"
	tree = recoverFromPreorder(traversal)
	tree.Floor()
}
