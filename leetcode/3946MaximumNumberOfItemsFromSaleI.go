package leetcode

// items = [[6,2],[2,6],[3,4]], budget = 9
func maximumSaleItems(items [][]int, budget int) int {
	f := make([]int, budget+1)
	mn := items[0][1]
	for _, item := range items {
		factor, price := item[0], item[1]
		mn = min(mn, price)
		cnt := 0
		for _, other := range items {
			if other[0]%factor == 0 {
				cnt++
			}
		}
		for j := budget; j >= price; j-- {
			f[j] = max(f[j], f[j-price]+cnt)
		}
	}
	ans := 0
	for j := 0; j <= budget; j++ {
		ans = max(ans, f[j]+(budget-j)/mn)
	}
	return ans
}
