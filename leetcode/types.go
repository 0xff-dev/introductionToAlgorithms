package leetcode

import (
	"fmt"
	"log"
)

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

func (t *TreeNode) Floor() {
	queue := []*TreeNode{t}
	layer := 1
	for len(queue) > 0 {
		fmt.Printf("layer: %d\n", layer)
		tmp := make([]*TreeNode, 0)
		for _, item := range queue {
			fmt.Print(item.Val, " ")
			if item.Left != nil {
				tmp = append(tmp, item.Left)
			}
			if item.Right != nil {
				tmp = append(tmp, item.Right)
			}
		}
		queue = tmp
		fmt.Println()
	}
}

type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}
