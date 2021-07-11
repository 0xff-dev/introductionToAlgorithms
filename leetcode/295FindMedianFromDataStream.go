package leetcode

const maxIdx = 100001

type MedianFinder struct {
	size     int
	positive []int
	negative []int
}

/** initialize your data structure here. */
func Constructor1() MedianFinder {
	return MedianFinder{
		size:     0,
		positive: make([]int, maxIdx),
		negative: make([]int, maxIdx),
	}
}

func (this *MedianFinder) AddNum(num int) {
	this.size++
	if num >= 0 {
		this.positive[num]++
		return
	}
	this.negative[-num]++
}

func (this *MedianFinder) FindMedian() float64 {
	leftNums := (this.size) / 2
	storeNums := 0
	pos := maxIdx - 1
	equalLeft := false
	// -2, -2, -1, 0, 2, 3, 5
	if this.size&1 == 1 {
		// -1, -1, 0, 1, 2
		// leftNums = 2
		for ; pos > 0; pos-- {
			if this.negative[pos] == 0 {
				continue
			}
			if equalLeft {
				return float64(-pos)
			}
			now := storeNums + this.negative[pos]
			if now == leftNums {
				equalLeft = true
				continue
			}
			if now > leftNums {
				return float64(-pos)
			}
			storeNums = now
		}
		for pos = 0; pos < maxIdx; pos++ {
			if this.positive[pos] == 0 {
				continue
			}
			if equalLeft {
				return float64(pos)
			}
			now := storeNums + this.positive[pos]
			if now == leftNums {
				equalLeft = true
				continue
			}
			if now > leftNums {
				return float64(pos)
			}
			storeNums = now
		}
		return 0.0
	}

	pre := maxIdx
	// -2,-1,0,0
	// -2, -2, -2, 1, 1, 3
	for pos = maxIdx - 1; pos > 0; pos-- {
		if this.negative[pos] == 0 {
			continue
		}
		if equalLeft {
			return float64(pre-pos) / 2
		}
		now := storeNums + this.negative[pos]
		if now == leftNums {
			pre = -pos
			equalLeft = true
			continue
		}
		if now > leftNums {
			return float64(pos)
		}
		pre = -pos
		storeNums = now
	}
	for pos = 0; pos < maxIdx; pos++ {
		if this.positive[pos] == 0 {
			continue
		}

		if equalLeft {
			return float64(pre+pos) / 2
		}
		now := storeNums + this.positive[pos]
		if now == leftNums {
			pre = pos
			equalLeft = true
			continue
		}
		if now > leftNums {
			return float64(pos)
		}
		pre = pos
		storeNums = now
	}
	return 0.0
}
