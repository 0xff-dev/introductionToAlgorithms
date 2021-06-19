package leetcode

import "testing"

func TestSumOfLeftLeaves(t *testing.T) {
    tree := &TreeNode{
        Val: 3,
        Left: &TreeNode{Val: 9},
        Right: &TreeNode{
            Val: 20,
            Left: &TreeNode{Val: 15},
            Right: &TreeNode{Val: 7},
        },
    }

    if sum := sumOfLeftLeaves(tree); sum != 24 {
        t.Fatalf("expect 24 get %d", sum)
    }

    tree = &TreeNode{Val: 1}
    if sum := sumOfLeftLeaves(tree); sum != 0 {
        t.Fatalf("expect 0 get %d", sum)
    }
}
