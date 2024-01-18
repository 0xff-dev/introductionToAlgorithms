package leetcode

import (
	"reflect"
	"testing"
)

func TestPrintTree(t *testing.T) {
	tree := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}
	exp := [][]string{
		{"", "1", ""}, {"2", "", ""},
	}
	if r := printTree(tree); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	tree = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{Val: 3},
	}
	exp = [][]string{
		{"", "", "", "1", "", "", ""},
		{"", "2", "", "", "", "3", ""},
		{"", "", "4", "", "", "", ""},
	}
	if r := printTree(tree); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	tree = &TreeNode{Val: 1}
	exp = [][]string{{"1"}}
	if r := printTree(tree); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
