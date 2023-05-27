package leetcode

import (
	"math"
)

func stoneGameIII(stoneValue []int) string {
	l := len(stoneValue)
	cache := make([]int, l)
	for i := 0; i < l; i++ {
		cache[i] = math.MinInt
	}
	var dpf func(int) int
	dpf = func(i int) int {
		if i >= l {
			return 0
		}
		if cache[i] != math.MinInt {
			return cache[i]
		}
		for j, s := 0, 0; j < 3 && i+j < l; j++ {
			s += stoneValue[i+j]
			if r := s + dpf(i+j+1); r > cache[i] {
				cache[i] = r
			}
		}
		return cache[i]
	}
	s := dpf(0)
	if s > 0 {
		return "Alice"
	} else if s < 0 {
		return "Bob"
	}
	return "Tie"
}
