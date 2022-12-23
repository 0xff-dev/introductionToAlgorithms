package leetcode

func maxProfit309(prices []int) int {

	// 持有，冷冻，卖出
	max := 0
	// a=持有, b=卖出，c=冷冻
	a, b, c := -prices[0], 0, 0
	for i := 1; i < len(prices); i++ {
		a1 := c - prices[i]
		if a > a1 {
			a1 = a
		}
		b1 := b
		if a+prices[i] > b1 {
			b1 = a + prices[i]
		}

		c1 := c
		if b > c1 {
			c1 = b
		}
		a, b, c = a1, b1, c1
		if a > max {
			max = a
		}
		if b > max {
			max = b
		}
		if c > max {
			max = c
		}
	}
	return max
}
