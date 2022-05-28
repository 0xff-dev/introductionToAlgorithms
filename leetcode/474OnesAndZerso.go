package leetcode

import (
	"fmt"
)

func bubbleSort(strs []string, store [][2]int, dp [][3]int, sortIdx int) {
	for c := 0; c < len(strs)-1; c++ {
		for cmp := 0; cmp < len(strs)-1-c; cmp++ {
			if store[cmp][sortIdx] < store[cmp+1][sortIdx] {
				store[cmp], store[cmp+1] = store[cmp+1], store[cmp]
				strs[cmp], strs[cmp+1] = strs[cmp+1], strs[cmp]
				dp[cmp], dp[cmp+1] = dp[cmp+1], dp[cmp]
			}
		}
	}
}
func findMaxForm1(strs []string, m int, n int) int {
	store := make([][2]int, len(strs))
	max := 0
	dp := make([][3]int, len(strs))
	for idx := range strs {
		bs := []byte(strs[idx])
		for _, b := range bs {
			store[idx][b-'0']++
		}
		if store[idx][0] <= m && store[idx][1] <= n {
			dp[idx][0], dp[idx][1], dp[idx][2] = store[idx][0], store[idx][1], 1
		}
	}
	sortIdx := 0
	if n > m {
		sortIdx = 1
	}
	bubbleSort(strs, store, dp, sortIdx)
	fmt.Printf("sort store: %+v\n", store)
	fmt.Printf("sort strs: %+v\n", strs)
	for idx := 1; idx < len(store); idx++ {
		for in := 0; in < idx; in++ {
			if zero, one := dp[in][0]+store[idx][0], dp[in][1]+store[idx][1]; zero <= m && one <= n && dp[in][2]+1 > dp[idx][2] {
				dp[idx][0], dp[idx][1] = zero, one
				dp[idx][2] = dp[in][2] + 1
				if dp[idx][2] > max {
					max = dp[idx][2]
				}
			}
		}
	}
	fmt.Printf("ans: %+v\n", dp)
	return max
}

// 假设：在前i中元素中，能够组成j个1和y0的最优解。
func findMaxForm(strs []string, m, n int) int {
	store := make([][2]int, len(strs))
	dp := make([][][]int, len(strs)+1)
	for x := 0; x < len(strs)+1; x++ {
		dp[x] = make([][]int, m+1)
		for y := 0; y < m+1; y++ {
			dp[x][y] = make([]int, n+1)
		}
		if x < len(strs) {
			for _, b := range []byte(strs[x]) {
				store[x][b-'0']++
			}
		}
	}

	max := 0
	for idx := 1; idx <= len(strs); idx++ {

		for zero := 0; zero <= m; zero++ {
			for one := 0; one <= n; one++ {
				t := dp[idx-1][zero][one]
				zeros, ones := store[idx-1][0], store[idx-1][1]

				if zero >= zeros && one >= ones {
					if r := dp[idx-1][zero-zeros][one-ones] + 1; r > t {
						t = r
					}
				}
				if t > dp[idx][zero][one] {
					dp[idx][zero][one] = t
				}
				if t > max {
					max = t
				}
			}
		}
	}
	return max
}
