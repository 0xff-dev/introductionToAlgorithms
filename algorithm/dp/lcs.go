package dp

import "fmt"
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func LCS(s1, s2 string) int {
	dp := map[string]int{}
	lenS1, lenS2 := len(s1), len(s2)
	for i := 0; i < lenS1; i++ {
		for j := 0; j < lenS2; j++ {
			key := fmt.Sprintf(keyFormat, i, j)
			key1, key2, key3 := fmt.Sprintf(keyFormat, i-1, j), fmt.Sprintf(keyFormat, i, j-1), fmt.Sprintf(keyFormat, i-1, j-1)
			if s1[i] == s2[j] {
				dp[key] = dp[key3] + 1
			} else {
				dp[key] = max(dp[key1], dp[key2])
			}
		}
	}
	printDp(dp, lenS1, lenS2)
	multiPath(dp, s1, s2, lenS1-1, lenS2-1, "")
	return dp[fmt.Sprintf(keyFormat, lenS1-1, lenS2-1)]
}

func printDp(dp map[string]int, x, y int) {
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			fmt.Print(dp[fmt.Sprintf(keyFormat, i, j)], " ")
		}
		fmt.Println()
	}
}

func multiPath(dp map[string]int, s1, s2 string ,x,  y int, now string)  {
	if x == -1 || y == -1 {
		fmt.Println("solution: ", now)
		return
	}
	key, key1, key2, key3 := fmt.Sprintf(keyFormat, x, y), fmt.Sprintf(keyFormat, x-1, y-1), fmt.Sprintf(keyFormat, x-1, y), fmt.Sprintf(keyFormat, x, y-1)
	if dp[key] == dp[key1]+1 && s1[x] == s2[y]{
		multiPath(dp, s1, s2, x-1, y-1, string(s1[x])+now)
		return
	}
	if dp[key] == dp[key2] {
		multiPath(dp, s1, s2, x-1, y, now)
	}
	if dp[key] == dp[key3] {
		multiPath(dp, s1, s2, x, y-1, now)
	}
}