package leetcode

func maxProfitIII(prices []int) int {
	min, max := prices[0], prices[len(prices)-1]
	ldp := make([]int, len(prices))
	lmax, rmax, result := 0, 0, 0
	for idx := 1; idx < len(prices); idx++ {
		ldp[idx] = lmax
		if prices[idx] < min {
			min = prices[idx]
			continue
		}
		if prices[idx]-min > lmax {
			lmax = prices[idx] - min
			ldp[idx] = lmax
		}
	}

	for idx := len(prices) - 2; idx >= 0; idx-- {
		if prices[idx] > max {
			max = prices[idx]
			continue
		}
		if r := max - prices[idx]; r > rmax {
			rmax = r
		}
		if r := rmax + ldp[idx]; r > result {
			result = r
		}
	}

	return result
}
