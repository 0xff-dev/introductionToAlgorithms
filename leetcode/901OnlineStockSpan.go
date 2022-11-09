package leetcode

type StockSpanner struct {
	less []int
	data []int
}

// 100 80 60 70 60 75 85
func Constructor901() StockSpanner {
	return StockSpanner{less: make([]int, 0), data: make([]int, 0)}
}

// 100 80 60 70 60 75 85
// 1   1  1  2  1  4
func (this *StockSpanner) Next(price int) int {
	this.less = append(this.less, 1)
	l := len(this.less)
	this.data = append(this.data, price)

	start := l - 2
	for ; start >= 0 && price >= this.data[start]; start -= this.less[start] {
		this.less[l-1] += this.less[start]
	}
	return this.less[len(this.less)-1]
}
