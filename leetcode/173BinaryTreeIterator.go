package leetcode

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor173(root *TreeNode) BSTIterator {
	r := BSTIterator{stack: make([]*TreeNode, 0)}
	w := root
	for ; w != nil; w = w.Left {
		r.stack = append(r.stack, w)
	}
	return r
}

func (this *BSTIterator) Next() int {
	l := len(this.stack)
	last := this.stack[l-1]
	this.stack = this.stack[:l-1]
	right := last.Right
	for ; right != nil; right = right.Left {
		this.stack = append(this.stack, right)
	}
	return last.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) != 0
}
