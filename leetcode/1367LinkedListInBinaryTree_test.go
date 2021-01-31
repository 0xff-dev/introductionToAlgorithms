package leetcode

import "testing"

func TestIsSubPath(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  4,
			Left: nil,
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   6,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
			},
			Right: nil,
		},
	}
	l := &ListNode{4, &ListNode{2, &ListNode{8, nil}}}
	if !isSubPath(l, tree) {
		t.Fatalf("is subset")
	}

	l = &ListNode{1, &ListNode{4, &ListNode{2, &ListNode{6, nil}}}}
	if !isSubPath(l, tree) {
		t.Fatalf("is subset")
	}

	l = &ListNode{1, &ListNode{4, &ListNode{2, &ListNode{6, &ListNode{8, nil}}}}}
	if isSubPath(l, tree) {
		t.Fatalf("isn't subset")
	}

	tree = &TreeNode{1, nil, nil}
	if isSubPath(l, tree) {
		t.Fatalf("isn't subset")
	}
}
