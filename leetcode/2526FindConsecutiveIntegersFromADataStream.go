package leetcode

type DataStream struct {
	value, k, index int

	top []int
	cnt map[int]int
}

func Constructor2526(value int, k int) DataStream {
	ds := DataStream{
		value: value,
		k:     k,
		top:   make([]int, k),
		cnt:   make(map[int]int),
		index: 0,
	}
	for i := range ds.top {
		ds.top[i] = -1
	}
	return ds
}

func (this *DataStream) Consec(num int) bool {
	if this.top[this.index] == -1 {
		this.top[this.index] = num
		this.cnt[num]++
		this.index = (this.index + 1) % this.k
		if this.index != 0 {
			return false
		}
	} else {
		remove := this.top[this.index]
		this.cnt[remove]--
		if this.cnt[remove] == 0 {
			delete(this.cnt, remove)
		}
		this.top[this.index] = num
		this.cnt[num]++
		this.index = (this.index + 1) % this.k
	}

	_, ok := this.cnt[this.value]
	return ok && len(this.cnt) == 1
}
