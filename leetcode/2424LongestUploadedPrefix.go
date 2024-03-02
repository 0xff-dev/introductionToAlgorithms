package leetcode

type LUPrefix struct {
	prefix []bool
	index  int
}

func Constructor2424(n int) LUPrefix {
	return LUPrefix{prefix: make([]bool, n+1), index: 0}
}

func (this *LUPrefix) Upload(video int) {
	this.prefix[video] = true
	next := this.index + 1
	for ; next < len(this.prefix) && this.prefix[next]; next++ {
	}
	this.index = next - 1
}

func (this *LUPrefix) Longest() int {
	return this.index
}
