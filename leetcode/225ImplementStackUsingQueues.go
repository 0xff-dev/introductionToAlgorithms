package leetcode

type MyStack struct {
	q1 []int
}

// 1, 2, 3
func Constructor225() MyStack {
	return MyStack{
		q1: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	tmp := make([]int, len(this.q1))
	// 先将queue清空，按顺序放到另一个queue中
	for i := 0; i < len(this.q1); i++ {
		tmp[i] = this.q1[i]
	}
	this.q1 = []int{x}
	// 原有元素入队
	this.q1 = append(this.q1, tmp...)
}

func (this *MyStack) Pop() int {
	x := this.q1[0]
	this.q1 = this.q1[1:]
	return x
}

func (this *MyStack) Top() int {
	return this.q1[0]
}

func (this *MyStack) Empty() bool {
	return len(this.q1) == 0
}
