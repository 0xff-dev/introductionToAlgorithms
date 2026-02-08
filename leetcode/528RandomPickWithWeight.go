package leetcode

import (
	"math/rand"
	"sort"
	"time"
)

type Solution528 struct {
	sum       int
	prefixSum []int

	r *rand.Rand
}

func Constructor528(w []int) Solution528 {
	sum := 0

	prefixSum := make([]int, len(w))
	for i := range w {
		sum += w[i]
		prefixSum[i] = sum
	}
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	return Solution528{
		sum:       sum,
		prefixSum: prefixSum,
		r:         seed,
	}
}

func (this *Solution528) PickIndex() int {
	ret := this.r.Intn(this.sum)
	return sort.Search(len(this.prefixSum), func(i int) bool {
		return this.prefixSum[i] > ret
	})
}
