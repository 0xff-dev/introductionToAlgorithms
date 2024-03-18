package leetcode

type FNode struct {
	Val        int
	Prev, Next *FNode
	Child      *FNode
}

func flatten430(root *FNode) *FNode {
	if root == nil {
		return nil
	}
	cur := &FNode{Val: root.Val}
	h := cur
	if root.Child != nil {
		if c := flatten430(root.Child); c != nil {
			cur.Next = c
			c.Prev = cur
		}

	}
	for ; cur.Next != nil; cur = cur.Next {

	}
	if c := flatten430(root.Next); c != nil {
		cur.Next = c
		c.Prev = cur
	}
	return h
}
