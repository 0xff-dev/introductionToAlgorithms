package challengeprogrammingdatastructure

import "testing"

func TestBinaryTreePrint(t *testing.T) {
	tree := []tree83{
		{0, 1, 4},
		{1, 2, 3},
		{2, -1, -1},
		{3, -1, -1},
		{4, 5, 8},
		{5, 6, 7},
		{6, -1, -1},
		{7, -1, -1},
		{8, -1, -1},
	}
	BinaryTreePritn(tree)
}