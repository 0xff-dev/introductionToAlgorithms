package offer

import "fmt"

type Node struct {
	Val int
	Next *Node
}

func PrintListReversing(head *Node) {
	if head == nil {
		return
	}
	if head.Next != nil {
		PrintListReversing(head.Next)
	}
	fmt.Println("now value: ", head.Val)
}
