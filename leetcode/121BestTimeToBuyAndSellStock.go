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
