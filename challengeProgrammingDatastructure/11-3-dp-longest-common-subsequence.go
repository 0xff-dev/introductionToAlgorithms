package challengeprogrammingdatastructure

import "fmt"

func findSubsequence(dp [][]int, x, y int, s1, s2, now string, ans *[]string) {
	if x <= 0 || y <= 0 {
		return
	}
	if s1[x-1] == s2[y-1] {
		now = fmt.Sprintf("%c", s1[x-1]) + now
		if x == 1 || y == 1 {
			*ans = append(*ans, now)
		} else {
			findSubsequence(dp, x-1, y-1, s1, s2, now, ans)
		}
		return
	}
	if dp[x][y] == dp[x-1][y] {
		findSubsequence(dp, x-1, y, s1, s2, now, ans)
	}
	if dp[x][y] == dp[x][y-1] {
		findSubsequence(dp, x, y-1, s1, s2, now, ans)
	}
}
func LongestCommonSubsequence(s1, s2 string) int {
	ls1, ls2 := len(s1), len(s2)
	dp := make([][]int, ls1+1)
	for i := 0; i <= ls1; i++ {
		dp[i] = make([]int, ls2+1)
	}
	for row := 1; row <= ls1; row++ {
		for col := 1; col <= ls2; col++ {
			if s1[row-1] == s2[col-1] {
				dp[row][col] = dp[row-1][col-1] + 1
				continue
			}
			dp[row][col] = dp[row-1][col]
			if dp[row][col-1] > dp[row][col] {
				dp[row][col] = dp[row][col-1]
			}
		}
	}
	ans := []string{}
	findSubsequence(dp, ls1, ls2, s1, s2, "", &ans)
	for _, s := range ans {
		fmt.Printf("find %s\n", s)
	}
	fmt.Println()
	return dp[ls1][ls2]
}
