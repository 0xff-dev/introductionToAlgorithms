package leetcode

import "testing"

func TestTree2Str(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 4},
		},
		Right: &TreeNode{Val: 3},
	}

	if r := tree2str(tree); r != "1(2(4))(3)" {
		t.Fatalf("expect '%s' get '%s'", "1(2(4))(3)", r)
	}

	tree = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{Val: 3},
	}

	if r := tree2str(tree); r != "1(2()(4))(3)" {
		t.Fatalf("expect '%s' get '%s'", "1(2()(4))(3)", r)
	}
}
