package offer

import "fmt"

func ConvertBinarySearchTree(tree *TreeNode) {
	var head *TreeNode
	auxBinarySearchTree(tree, &head)
	for walker := head; walker != nil; walker = walker.Left {
		fmt.Println("data: ", walker.Val)
	}
}

func auxBinarySearchTree(tree *TreeNode, head **TreeNode) {
	if tree == nil {
		return
	}
	currNode := tree
	if tree.Left != nil {
		auxBinarySearchTree(tree.Left, head)
	}
	currNode.Left = *head
	if *head != nil {
		(*head).Right = currNode
	}
	*head = currNode
	if currNode.Right != nil {
		auxBinarySearchTree(tree.Right, head)
	}
}
