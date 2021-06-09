package leetcode


const inf = 0xffff

func palindromePartition(s string, k int) int {
    length := len(s)
    dp := make([][]int, length+1)
    for idx := 0; idx <= length; idx++ {
        dp[idx] = make([]int, k+1)
        for col := 0; col <= k; col++ {
            dp[idx][col] = inf
        }
    }

    for charIdx := 1; charIdx <= length; charIdx++ {
        minSplit := charIdx
        if minSplit > k {
            minSplit = k
        }
        for split := 1; split <= minSplit; split ++ {
            if split == 1 {
                count := steps(s[:charIdx])
                if dp[charIdx][split] > count {
                    dp[charIdx][split] = count
                }
                continue
            } 
            for separator := charIdx; separator >= split; separator-- {
                tmp := dp[separator-1][split-1]+steps(s[separator-1:charIdx])
                if dp[charIdx][split] > tmp {
                    dp[charIdx][split] = tmp
                }
            }
        }
    }
    return dp[length][k]
}

func steps(str string) int {
    length := len(str)
    if length == 0 || length == 1 {
        return 0
    }
    step := 0
    for s, e := 0, length-1; s < e; s, e = s+1, e-1 {
        if str[s] != str[e] {
            step++
        }
    }
    return step
}
