package dp

import "fmt"

var (
	price = []int{1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
)

// 10 个为例子, 输出他的解
func MaxPrice(steelBar int) int {
	if steelBar == 0 {
		return 0
	}
	dp := make([]int, steelBar+1)
	dp[1] = 1
	solution := make([]int, steelBar+1)
	for index := 2; index <= steelBar; index++ {
		for inner := 1; inner <= index; inner++ {
			if price[inner-1]+dp[index-inner] > dp[index] {
				dp[index] = price[inner-1] + dp[index-inner]
				solution[index] = inner
			}
		}
	}
	for index := steelBar; index > 1; index = index - solution[index] {
		fmt.Println("solution: ", solution[index])
	}
	return dp[steelBar]
}
