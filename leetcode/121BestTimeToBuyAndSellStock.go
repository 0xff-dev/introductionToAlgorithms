package leetcode

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	res := 0
	max := prices[len(prices)-1]
	for idx := len(prices) - 2; idx >= 0; idx-- {
		tmp := max - prices[idx]
		if tmp > res {
			res = tmp
		}
		if tmp < 0 {
			max = prices[idx]
		}
	}
	return res
}

func maxProfit2(prices []int) int {
	ans := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
			continue
		}
		diff := prices[i] - min
		if diff > ans {
			ans = diff
		}
	}
	return ans
}
