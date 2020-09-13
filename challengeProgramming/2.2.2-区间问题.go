package challengeProgramming

/*
	n项工作，每一项的开始和结束时间分别是si, ti, 对于每一项工作，你都可以选择参与和不参与
	如果参与了，必须是全程参与，参与的工作不能重复。求出最多可以参与多少工作
*/

// 按照结束时间排序
func workEndTimeSort(start, end []int) {
	for cmp := 0; cmp < len(end)-1; cmp++ {
		for walker := 0; walker < len(end)-1-cmp; walker++ {
			if end[walker] > end[walker+1] {
				end[walker], end[walker+1] = end[walker+1], end[walker]
				start[walker], start[walker+1] = start[walker+1], start[walker]
			}
		}
	}
}

func MostInvolvedWork(start, end []int) int {
	if len(start) != len(end) || len(start) == 0 {
		return 0
	}
	workEndTimeSort(start, end)
	cnt := 1
	endTime := end[0]
	for index := 1; index < len(end); index++ {
		if start[index] >= endTime {
			cnt++
			endTime = end[index]
		}
	}
	return cnt
}
