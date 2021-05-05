package dp

import "fmt"

const minCost = -1

func min(a, b int) int {
    if a == -1 {
        return b
    }
    if a > b {
        return b
    }
    return a
}
//   有2，5，7三种硬币，无限制，用最少的数量找零钱
// 找状态
func minCoins(target int) int {
    if target == 0 {
        return 0
    }
    if target == 1 {
        // 无解
        return minCost
    }
	r := make([]int, target+1)
    for idx := 1; idx < target+1; idx++ {
        r[idx] = minCost
    }

    for i := 2; i <= target; i++ {
        now := minCost
        if r[i-2] != minCost {
            now = min(now, r[i-2]+1)
        }
        if i-5 >= 0 && r[i-5] != minCost {
            now = min(now, r[i-5]+1)
        }
        if i-7 >= 0 && r[i-7] != minCost {
            now = min(now, r[i-7]+1)
        }
        r[i] = now
    }
    for i := 0; i <= target; i++ {
        fmt.Printf("r[%d]=%d\n", i, r[i])
    }
	return r[target]
}
