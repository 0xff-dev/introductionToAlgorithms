package leetcode

import (
	"sort"
)

func minimumEffort(tasks [][]int) int {
	// 3+2+9+11+7
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1]-tasks[i][0] < tasks[j][1]-tasks[j][0]
	})

	ret := 0
	for i := range tasks {
		if ret+tasks[i][0] > tasks[i][1] {
			ret += tasks[i][0]
			continue
		}
		ret = tasks[i][1]
	}
	return ret
}
