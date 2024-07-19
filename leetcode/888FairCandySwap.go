package leetcode

import (
	"fmt"
	"sort"
)

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	sort.Ints(bobSizes)

	a, b := 0, 0
	for i := 0; i < len(aliceSizes); i++ {
		a += aliceSizes[i]
	}
	for i := 0; i < len(bobSizes); i++ {
		b += bobSizes[i]
	}

	for i := 0; i < len(aliceSizes); i++ {
		out := a - aliceSizes[i]
		idx := sort.Search(len(bobSizes), func(j int) bool {
			return out+bobSizes[j] >= b-bobSizes[j]+aliceSizes[i]
		})
		if idx != len(bobSizes) && out+bobSizes[idx] == b-bobSizes[idx]+aliceSizes[i] {
			return []int{aliceSizes[i], bobSizes[idx]}
		}
	}
	return nil
}
