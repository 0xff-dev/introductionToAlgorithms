package leetcode

func numDistinct(s string, t string) int {
	lens, lent := len(s), len(t)
	dp := make([][]int, 2)
	for idx := 0; idx < 2; idx++ {
		dp[idx] = make([]int, lent+1)
		dp[idx][0] = 1
	}

	// ''=''
	loop := 1
	for row := 0; row < lens; row++ {
		for col := 1; col <= lent; col++ {
			dp[loop][col] = dp[1-loop][col]
			if s[row] == t[col-1] {
				dp[loop][col] += dp[1-loop][col-1]
			}
		}
		loop = 1 - loop
	}
	/*
	   for row := 1; row <= lens; row++ {
	       for col := 1; col <= lent; col++ {
	           dp[row][col] = dp[row-1][col]
	           if s[row-1] == t[col-1] {
	               dp[row][col] += dp[row-1][col-1]
	           }
	       }
	   }
	*/
	return dp[1-loop][lent]
}

func numDistinct1(s, t string) int {
	ls, lt := len(s), len(t)
	dp := make([]int, lt+1)
	dp[0] = 1
	var diag, tmp int
	for i := range ls {
		diag = dp[0]
		for j := 1; j <= lt; j++ {
			tmp, diag = diag, dp[j]
			diag = dp[j]
			if s[i] == t[j-1] {
				dp[j] += tmp
			}
		}
	}
	return dp[lt]
}
