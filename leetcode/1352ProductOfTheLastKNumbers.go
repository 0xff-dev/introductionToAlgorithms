package leetcode

type ProductOfNumbers struct {
	data  []int
	zeros map[int]struct{}
}

func Constructor1352() ProductOfNumbers {
	return ProductOfNumbers{data: make([]int, 0), zeros: map[int]struct{}{}}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.data = append(this.data, num)
		this.zeros[len(this.data)-1] = struct{}{}
		return
	}
	if len(this.data) != 0 {
		if last := this.data[len(this.data)-1]; last != 0 {
			num *= last
		}
	}
	this.data = append(this.data, num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	end := len(this.data) - 1
	start := end - k
	for i := range this.zeros {
		if i > start && i <= end {
			return 0
		}
	}
	if start < 0 || this.data[start] == 0 {
		return this.data[end]
	}
	return this.data[end] / this.data[start]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
