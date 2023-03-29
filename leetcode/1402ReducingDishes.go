package leetcode

import (
	"sort"
)

func maxSatisfaction(satisfaction []int) int {
	sort.Ints(satisfaction)
	ans := 0
	x := 0
	// 好的算法总是在不经意间让你醍醐灌顶
	for i := len(satisfaction) - 1; i >= 0 && x+satisfaction[i] > 0; i-- {
		x += satisfaction[i]
		ans += x
	}
	return ans
	// sort.Slice(satisfaction, func(i, j int) bool {
	// 	if satisfaction[i] < 0 && satisfaction[j] < 0 {
	// 		return -satisfaction[i] < -satisfaction[j]
	// 	}
	// 	if satisfaction[i] < 0 {
	// 		return true
	// 	}
	// 	if satisfaction[j] < 0 {
	// 		return false
	// 	}
	// 	return satisfaction[i] < satisfaction[j]
	// })
	// // positiveSum = a+b+c
	// lastIndex, ans, positiveSum := -1, 0, 0
	// negativeSum := 0

	// // base = a*0 + b*1 + c*2
	// base, costTime := 0, 1
	// for i := 0; i < len(satisfaction); i++ {
	// 	if satisfaction[i] < 0 {
	// 		negativeSum += (i + 1) * satisfaction[i]
	// 		satisfaction[i] = negativeSum
	// 		lastIndex = i
	// 		continue
	// 	}
	// 	positiveSum += satisfaction[i]

	// 	base += satisfaction[i] * costTime
	// 	costTime++
	// }
	// if lastIndex == len(satisfaction)-1 {
	// 	return ans
	// }
	// // 不选择负数的时候
	// ans = base
	// if lastIndex == -1 {
	// 	return ans
	// }
	// // -1 -2 -3 0 3 5
	// for negative := 0; negative <= lastIndex; negative++ {
	// 	r := base + (negative+1)*positiveSum + satisfaction[negative]
	// 	if r > ans {
	// 		ans = r
	// 	}
	// }
	// return ans
	// var dfs func(int, int, int)
	// ans := 0
	// dfs = func(index, costTime, sum int) {
	// 	if sum > ans {
	// 		ans = sum
	// 	}
	// 	if index >= len(satisfaction) {
	// 		return
	// 	}
	// 	// 选择做
	// 	dfs(index+1, costTime+1, sum+(costTime+1)*satisfaction[index])
	// 	//丢弃
	// 	dfs(index+1, costTime, sum)
	// }
	// dfs(0, 0, 0)
	// return ans
}
