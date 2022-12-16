package leetcode

// d, e, f,
// c, b
type MyQueue struct {
	stack1, stack2 []int
}

func Constructor232() MyQueue {
	return MyQueue{
		stack1: make([]int, 0),
		stack2: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.stack1 = append(this.stack1, x)
}

func (this *MyQueue) Pop() int {
	if l := len(this.stack2); l > 0 {
		x := this.stack2[l-1]
		this.stack2 = this.stack2[:l-1]
		return x
	}
	// a, b, c
	l := len(this.stack1)
	for i := l - 1; i > 0; i-- {
		this.stack2 = append(this.stack2, this.stack1[i])
	}
	x := this.stack1[0]
	this.stack1 = this.stack1[:0]
	return x
}

func (this *MyQueue) Peek() int {
	if l := len(this.stack2); l > 0 {
		return this.stack2[l-1]
	}
	l := len(this.stack1)
	for i := l - 1; i >= 0; i-- {
		this.stack2 = append(this.stack2, this.stack1[i])
	}
	this.stack1 = this.stack1[:0]
	return this.stack2[len(this.stack2)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.stack1) == 0 && len(this.stack2) == 0
}
