package leetcode

type DataStream struct {
	value, k, cnt int
}

func Constructor2526(value int, k int) DataStream {
	return DataStream{
		value: value,
		k:     k,
		cnt:   0,
	}
}

// value = 1, k = 2
// 1, 2, 1, 1, 1

func (this *DataStream) Consec(num int) bool {
	if num == this.value {
		this.cnt++
		return this.cnt >= this.k
	}
	this.cnt = 0
	return false
}

/**
 * Your DataStream object will be instantiated and called as such:
 * obj := Constructor(value, k);
 * param_1 := obj.Consec(num);
 */
