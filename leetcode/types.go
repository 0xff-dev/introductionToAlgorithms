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

type MultiChildNode struct {
	Val      int
	Children []*MultiChildNode
}

/*
Trie Node
*/

// TrieNOde support low letter
type TrieNode struct {
	// a- z 0- 25
	Child [26]*TrieNode
	Count uint8
}

// ab
// abc
func (tn *TrieNode) Insert(str string) {
	root := tn
	for _, bt := range []byte(str) {
		if root.Child[bt-'a'] == nil {
			root.Child[bt-'a'] = &TrieNode{Count: 0, Child: [26]*TrieNode{}}
			root.Count++
		}
		root = root.Child[bt-'a']
	}
}

func (tn *TrieNode) Search(pattern string) bool {
	root := tn
	for _, bt := range []byte(pattern) {
		if root == nil || root.Child[bt-'a'] == nil {
			return false
		}
		root = root.Child[bt-'a']
	}
	return true
}

type QuadNode struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *QuadNode
	TopRight    *QuadNode
	BottomLeft  *QuadNode
	BottomRight *QuadNode
}
