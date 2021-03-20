package leetcode

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		intervals = append(intervals, newInterval)
		return intervals
	}
	if len(newInterval) == 0 {
		return intervals
	}
	r := make([][]int, 0)
	beginNum, endNum := newInterval[0], newInterval[1]
	idx := 0
	for ; idx < len(intervals); idx++ {
		if beginNum > intervals[idx][1] {
			r = append(r, intervals[idx])
			continue
		}
		if beginNum > intervals[idx][0] {
			beginNum = intervals[idx][0]
		}
		break
	}
	found := false

	for i := idx; i < len(intervals); i++ {
		if endNum > intervals[i][1] {
			continue
		}

		if endNum >= intervals[i][0] {
			endNum = intervals[i][1]
			i++
		}
		r = append(r, []int{beginNum, endNum})
		r = append(r, intervals[i:]...)
		found = true
		break
	}
	if !found {
		r = append(r, []int{beginNum, newInterval[1]})
	}
	return r
}
