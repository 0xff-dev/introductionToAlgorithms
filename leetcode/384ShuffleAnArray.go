package leetcode

import (
	"math/rand"
	"time"
)

type Solution384 struct {
	origin []int
	r      *rand.Rand
}

func Constructor384(nums []int) Solution384 {
	return Solution384{
		origin: nums,
		r:      rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (this *Solution384) Reset() []int {
	this.r = rand.New(rand.NewSource(time.Now().UnixNano()))
	return this.origin
}

func (this *Solution384) Shuffle() []int {
	n := len(this.origin)
	cur := make([]int, len(this.origin))
	copy(cur, this.origin)
	for i := n - 1; i > 0; i-- {
		j := this.r.Intn(i + 1)
		cur[i], cur[j] = cur[j], cur[i]
	}
	return cur
}
