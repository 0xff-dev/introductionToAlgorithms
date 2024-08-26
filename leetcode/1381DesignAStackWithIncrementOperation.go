package leetcode

type CustomStack struct {
	index int
	stack []int
}

func Constructor1381(maxSize int) CustomStack {
	return CustomStack{
		stack: make([]int, maxSize),
		index: -1,
	}
}

func (this *CustomStack) Push(x int) {
	if this.index == len(this.stack)-1 {
		return
	}

	this.index++
	this.stack[this.index] = x
}

func (this *CustomStack) Pop() int {
	if this.index == -1 {
		return -1
	}

	top := this.stack[this.index]
	this.index--
	return top
}

func (this *CustomStack) Increment(k int, val int) {
	if this.index == -1 {
		return
	}

	// 1, 2, 3, 4,
	// 3
	for i := 0; i <= this.index && i < k; i++ {
		this.stack[i] += val
	}
}
