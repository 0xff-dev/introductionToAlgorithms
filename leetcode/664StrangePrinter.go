package leetcode

func strangePrinter(s string) int {
	// a,a,a
	// a,a,b,
	// a,a,b,a,b

	// a,b,a -> 2
	// aa,b -> 2
	// a, b,b -> 2
	// a, b,c -> 3

	l := len(s)
	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
		for j := 0; j < l; j++ {
			dp[i][j] = l
		}
	}

	for ll := 1; ll <= l; ll++ {
		for start := 0; start <= l-ll; start++ {
			end := start + ll - 1
			j := -1
			for i := start; i < end; i++ {
				if s[i] != s[end] && j == -1 {
					j = i
				}
				if j != -1 {
					if x := dp[j][i] + dp[i+1][end] + 1; x < dp[start][end] {
						dp[start][end] = x
					}
				}
			}
			if j == -1 {
				dp[start][end] = 0
			}
		}
	}
	return dp[0][l-1] + 1
	/*
		for ll := 2; ll <= l; ll++ {
			for i := 0; i <= l-ll; i++ {
				end := i + ll - 1
				for k := end; k > i; k-- {
					total := store[end-1]
					if k == end {
						add := dp[end][end]
						if s[end] == s[end-1] {
							add = 0
						}
						if dp[i][end] == 0 || dp[i][k-1]+add < dp[i][end] {
							dp[i][end] = dp[i][k-1] + add
						}
						continue
					}

					total[s[end]-'a'] -= store[k-1][s[end]-'a']
					add := 1
					if total[s[end]-'a'] > 0 {
						//存在
						add = 0
					}
					if dp[i][end] == 0 || dp[i][k-1]+dp[k][end-1]+add < dp[i][end] {
						dp[i][end] = dp[i][k-1] + dp[k][end-1] + add
					}
				}
			}
		}
	*/
}
