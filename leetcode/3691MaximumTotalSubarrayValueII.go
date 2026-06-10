package leetcode

import (
	"container/heap"
	"math/bits"
)

func maxTotalValue3691(nums []int, k int) int64 {
	n := len(nums)
	logn := bits.Len(uint(n))
	stMax := make([][]int, n)
	stMin := make([][]int, n)
	for i := range stMax {
		stMax[i] = make([]int, logn)
		stMin[i] = make([]int, logn)
		stMax[i][0] = nums[i]
		stMin[i][0] = nums[i]
	}
	for j := 1; j < logn; j++ {
		for i := 0; i+(1<<j) <= n; i++ {
			stMax[i][j] = max(stMax[i][j-1], stMax[i+(1<<(j-1))][j-1])
			stMin[i][j] = min(stMin[i][j-1], stMin[i+(1<<(j-1))][j-1])
		}
	}
	queryMax := func(l, r int) int {
		j := bits.Len(uint(r-l+1)) - 1
		return max(stMax[l][j], stMax[r-(1<<j)+1][j])
	}
	queryMin := func(l, r int) int {
		j := bits.Len(uint(r-l+1)) - 1
		return min(stMin[l][j], stMin[r-(1<<j)+1][j])
	}
	h := &hp{}
	for l := 0; l < n; l++ {
		heap.Push(h, tuple{queryMax(l, n-1) - queryMin(l, n-1), l, n - 1})
	}
	var ans int64 = 0
	for ; k > 0; k-- {
		t := heap.Pop(h).(tuple)
		ans += int64(t.val)
		if t.r > t.l {
			heap.Push(h, tuple{queryMax(t.l, t.r-1) - queryMin(t.l, t.r-1), t.l, t.r - 1})
		}
	}
	return ans
}

type tuple struct{ val, l, r int }
type hp []tuple

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].val > h[j].val }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
