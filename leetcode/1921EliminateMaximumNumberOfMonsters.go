package leetcode

import (
	"sort"
)

func eliminateMaximum(dist []int, speed []int) int {
	minutes := make([]float64, len(dist))
	for i := 0; i < len(dist); i++ {
		minutes[i] = float64(dist[i]) / float64(speed[i])
	}
	sort.Float64s(minutes)
	cost := 0
	ans := 0
	full := true
	for i := 0; i < len(minutes); i++ {
		if full {
			ans++
			full = false
			continue
		}
		// 抢需要充电，看看下一个怪物到城市是否小于等于1分钟，如果是，就是没有时间充电了，如果不是
		leftMinutes := minutes[i] - float64(cost)
		if leftMinutes <= 1.0 {
			break
		}
		i--
		full = true
		cost++
	}
	return ans
}
