package leetcode

import "slices"

func maximumSaleItems3947(items [][]int, budget int) int {
	slices.SortFunc(items, func(a, b []int) int {
		return a[1] - b[1]
	})
	mx := 0
	for _, item := range items {
		mx = max(mx, item[0])
	}
	freq := make([]int, mx+1)
	for _, item := range items {
		freq[item[0]]++
	}
	for i := 1; i <= mx; i++ {
		if freq[i] > 0 {
			for j := i * 2; j <= mx; j += i {
				freq[i] += freq[j]
			}
		}
	}

	mn := items[0][1]
	ans := 0
	for _, item := range items {
		factor, price := item[0], item[1]
		if price > mn*2 {
			break
		}
		cnt := freq[factor] - 1
		take := min(cnt, budget/price)
		ans += take * 2
		budget -= take * price
	}
	ans += budget / mn
	return ans
}
