package challengeprogrammingdatastructure

import "testing"

func TestRootedTree(t *testing.T) {
	tree := []treeDefine{
		{0, 3, []int{1, 4, 10}},
		{1, 2, []int{2, 3}},
		{2, 0, nil},
		{3, 0, nil},
		{4, 3, []int{5, 6, 7}},
		{5, 0, nil},
		{6, 0, nil},
		{7, 2, []int{8, 9}},
		{8, 0, nil},
		{9, 0, nil},
		{10, 2, []int{11, 12}},
		{11, 0, nil},
		{12, 0, nil},
	}
	RootedTree(tree)
}
