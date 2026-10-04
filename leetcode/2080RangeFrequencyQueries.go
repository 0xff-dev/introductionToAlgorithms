package leetcode

type RangeFreqQuery struct {
	freq map[int][]int
}

func Constructor2080(arr []int) RangeFreqQuery {
	r := RangeFreqQuery{
		freq: make(map[int][]int),
	}
	for i := range arr {
		r.freq[arr[i]] = append(r.freq[arr[i]], i)
	}
	return r
}

func search2080(n int, ok func(int) bool) int {
	// return indies[mid] >= left
	// return indies[mmid] > right
	left, right := 0, n
	for left < right {
		mid := left + (right-left)/2
		if ok(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	indies := this.freq[value]
	a := search2080(len(indies), func(mid int) bool {
		return indies[mid] >= left
	})
	b := search2080(len(indies), func(mid int) bool {
		return indies[mid] > right
	})
	return b - a
}

/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */
