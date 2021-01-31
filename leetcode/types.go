package leetcode

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Dis() {
	for walker := l; walker != nil; walker = walker.Next {
		log.Printf("val is [%d]", walker.Val)
	}
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func (n *Node) Dis() {
	for walker := n; walker != nil; walker = walker.Next {
		log.Printf("val is [%d]", walker.Val)
	}
}

// tree node define
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
