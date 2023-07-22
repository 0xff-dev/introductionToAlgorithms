package leetcode

import "sort"

type pair1338 struct {
	v, c int
}

func minSetSize(arr []int) int {
	pairs := make([]pair1338, 0)
	exist := make(map[int]int)
	l := len(arr)
	for idx := 0; idx < l; idx++ {
		index, ok := exist[arr[idx]]
		if !ok {
			pairs = append(pairs, pair1338{v: arr[idx], c: 1})
			exist[arr[idx]] = len(pairs) - 1
			continue
		}
		pairs[index].c++
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].c > pairs[j].c
	})
	sum := 0
	ans := 0
	for _, pair := range pairs {
		sum += pair.c
		ans++
		if sum >= l/2 {
			break
		}
	}

	return ans
}
