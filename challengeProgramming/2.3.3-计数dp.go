package challengeProgramming


/*
   有n个物品，划分为m组，求划分方法模M的余数
   计数问题中重复问题一定要注意
*/

func DivideN2MGroups(n, m, M int) int {
    dp := make([][]int, m+1)
    for i := 0; i < m+1; i++ {
        dp[i] = make([]int, n+1)
    }
    dp[0][0] = 1
    for row := 1; row <= m; row++ {
        for col := 0; col <= n; col++ {
            if col - row >= 0 {
                dp[row][col] = (dp[row][col-row]+dp[row-1][col]) % M
            } else {
                dp[row][col] = dp[row-1][col]
            }
        }
    }
    return dp[m][n]
}
