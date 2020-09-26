package challengeProgramming

import "fmt"

/*
01背包问题
*/

func Backpack(weights, values []int, capacity int) int {
	if len(weights) != len(values) {
		return -1
	}
	dp := make([]int, capacity+1)
	for index, w := range weights {
		for v := capacity; v >= 0; v-- {
			if v >= w {
				dp[v] = max(dp[v], dp[v-w]+values[index])
			}
		}
	}
	return dp[capacity]
}

/*
最长公共子序列问题
*/
func LongestCommonSubsequence(s1, s2 string) int {
	lenS1, lenS2 := len(s1), len(s2)
	if lenS1 == 0 || lenS2 == 0 {
		return 0
	}
	dp := make([][]int, lenS2+1)
	for i := 0; i < lenS2+1; i++ {
		dp[i] = make([]int, lenS1+1)
	}
	for col := 1; col < lenS1+1; col++ {
		for row := 1; row < lenS2+1; row++ {
			if s1[col-1] == s2[row-1] {
				dp[row][col] = dp[row-1][col-1] + 1
			} else {
				dp[row][col] = max(dp[row-1][col], dp[row][col-1])
			}
		}
	}
	findPath(dp, s2, s1, lenS2, lenS1, "")
	return dp[lenS2][lenS1]
}

func findPath(dp [][]int, s1, s2 string, x, y int, ans string) {
	if x == 0 || y == 0 {
		fmt.Println("ans: ", ans)
		return
	}
	cx, cy := s1[x-1], s2[y-1]
	if cx == cy {
		findPath(dp, s1, s2, x-1, y-1, string(cx)+ans)
		return
	}
	if dp[x][y] == dp[x-1][y] {
		findPath(dp, s1, s2, x-1, y, ans)
	}
	if dp[x][y] == dp[x][y-1] {
		findPath(dp, s1, s2, x, y-1, ans)
	}
}

// 明天实现
func LCSbyScrollingArray(s1, s2 string) string {
	return ""
}
