package leetcode

type pair714 struct {
	start, end int
}

// Not A good solution, but it was accepted.
func maxProfit714(prices []int, fee int) int {
	start, end := 0, 1
	pairs := make([]pair714, 0)
	for ; end < len(prices); end++ {
		if prices[end] >= prices[end-1] {
			continue
		}
		if prices[end-1] > prices[start] {
			pairs = append(pairs, pair714{start, end - 1})
		}
		start = end
	}
	if prices[end-1] > prices[start] {
		pairs = append(pairs, pair714{start, end - 1})
	}
	ans := 0
	if len(pairs) == 0 {
		return ans
	}

	dp := make([]int, len(pairs))

	if r := prices[pairs[0].end] - prices[pairs[0].start] - fee; r > 0 {
		dp[0] = r
	}
	if dp[0] > ans {
		ans = dp[0]
	}
	for i := 1; i < len(pairs); i++ {
		for j := i; j >= 0; j-- {
			diff := prices[pairs[i].end] - prices[pairs[j].start] - fee
			left := 0
			if j > 0 {
				left = dp[j-1]
			}

			if diff > 0 {
				left += diff
			}
			if left > dp[i] {
				dp[i] = left
			}
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

// Editorial
func maxProfit714_1(prices []int, fee int) int {
	free := make([]int, len(prices))
	hold := make([]int, len(prices))
	hold[0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		if r := hold[i-1] + prices[i] - fee; r > free[i] {
			free[i] = r
		}
		if r := free[i] - prices[i]; r > hold[i] {
			hold[i] = r
		}
	}
	return free[len(prices)-1]
}
