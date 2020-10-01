package challengeProgramming

import "fmt"

// 完全背包, 某个物品可以无限制的选择
func CompleteBackpack(weights, values []int, capacity int) int {
	lw, lv := len(weights), len(values)
	if lw != lv {
		fmt.Println("weights does't equal values")
		return 0
	}

	dp := make([]int, capacity+1)
	for goods := 0; goods < lw; goods++ {
		for _cap := capacity; _cap > 0; _cap-- {
			// 支持选择同一个物品多个, 0-n
			for cnt := 1; cnt*weights[goods] <= _cap; cnt++ {
				dp[_cap] = max(dp[_cap], dp[_cap-weights[goods]*cnt]+values[goods]*cnt)
			}
		}
	}
	return dp[capacity]
}

/*
   有n个不同大小的数字，每个有m个，求选择他们中的若干是否可以是他们的和为K
   这道题可以考虑dp， 母函数
*/

// 母函数思想，可以这么理解，选择一个的时候可以得到那些值，在选择第二个，可以得到另一波的值。这样
// 根据系数即可得到目标的有多少中组合
func GeneratingFunction(values, nums []int, k int) int {
	// 一个基本的生成式
	if len(values) == 0 {
		return 0
	}
	// coefficient就是系数表达式, tmp 辅助计算
	// 0+0*x^1+0*x^2+......
	coefficient, tmp := make([]int, k+1), make([]int, k+1)
	for c, v := 0, 0; c <= nums[0] && v <= k; c, v = c+1, v+values[0] {
		coefficient[v] = 1
	}

	for index := 1; index < len(values); index++ {
		for idx := 0; idx <= k; idx++ {
			for c, v := 0, 0; c <= nums[index] && v+idx <= k; c, v = c+1, v+values[index] {
				tmp[v+idx] += coefficient[idx]
			}
		}
		for i := 0; i <= k; i++ {
			coefficient[i], tmp[i] = tmp[i], 0
		}
	}

	for index := 0; index < k+1; index++ {
		fmt.Printf("%dx^%d+ ", coefficient[index], index)
	}
	fmt.Println()
	return coefficient[k]
}
