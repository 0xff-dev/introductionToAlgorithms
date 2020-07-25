package offer

import "testing"

func TestIsChildren(t *testing.T) {
	tree1 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{1, nil, nil},
		Right: &TreeNode{3, nil, nil},
	}
	tree2 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	if IsChildren(tree1, tree2) == true {
		t.Fatalf("except false")
	}
}
