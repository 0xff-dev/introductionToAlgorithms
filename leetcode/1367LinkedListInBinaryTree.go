package leetcode

func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil && head != nil {
		return false
	}
	if head == nil {
		return true
	}
	found := false
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		front := queue[0]
		walker := head
		aux(front, walker, &found)
		if found {
			break
		}
		if front.Left != nil {
			queue = append(queue, front.Left)
		}
		if front.Right != nil {
			queue = append(queue, front.Right)
		}
		queue = queue[1:]
	}
	return found
}

func aux(treeRoot *TreeNode, list *ListNode, res *bool) {
	if *res {
		return
	}
	if treeRoot == nil && list != nil {
		return
	}
	if list == nil {
		*res = true
		return
	}
	if treeRoot != nil {
		if treeRoot.Val == list.Val {
			aux(treeRoot.Left, list.Next, res)
			aux(treeRoot.Right, list.Next, res)
		}
	}
}
