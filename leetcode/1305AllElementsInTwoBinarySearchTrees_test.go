package leetcode

import (
	"reflect"
	"testing"
)

func TestGetAllelements(t *testing.T) {
	t1 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 4},
	}
	t2 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 0},
		Right: &TreeNode{Val: 3},
	}
	exp := []int{0, 1, 1, 2, 3, 4}
	if r := getAllElements(t1, t2); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	t1 = &TreeNode{Val: 1, Right: &TreeNode{Val: 8}}
	t2 = &TreeNode{Val: 8, Left: &TreeNode{Val: 1}}
	exp = []int{1, 1, 8, 8}
	if r := getAllElements(t1, t2); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
