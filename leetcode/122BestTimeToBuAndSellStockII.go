package leetcode

func maxProfitII(prices []int) int {
	max := 0
	for index := 1; index < len(prices); index++ {
		if prices[index] > prices[index-1] {
			max += prices[index] - prices[index-1]
		}
	}
	return max
}
