package leetcode

import "sort"

type SegmentTree3501 struct {
	n   int
	arr []int
	seg []int
}

func NewSegmentTree3501(arr []int) *SegmentTree3501 {
	st := &SegmentTree3501{
		arr: arr,
		n:   len(arr),
		seg: make([]int, len(arr)*4),
	}
	st.build(1, 0, st.n-1)

	return st
}

func (st *SegmentTree3501) build(p, l, r int) {
	if l == r {
		st.seg[p] = st.arr[l]
		return
	}

	mid := (l + r) >> 1
	st.build(p<<1, l, mid)
	st.build(p<<1|1, mid+1, r)

	st.seg[p] = max(
		st.seg[p<<1],
		st.seg[p<<1|1],
	)
}

func (st *SegmentTree3501) queryInternal(p, l, r, L, R int) int {
	if L <= l && r <= R {
		return st.seg[p]
	}

	mid := (l + r) >> 1
	res := 0
	if L <= mid {
		res = max(res, st.queryInternal(p<<1, l, mid, L, R))
	}
	if R > mid {
		res = max(res, st.queryInternal(p<<1|1, mid+1, r, L, R))
	}

	return res
}

func (st *SegmentTree3501) Query(L, R int) int {
	if L > R {
		return 0
	}
	return st.queryInternal(1, 0, st.n-1, L, R)
}

func maxActiveSectionsAfterTrade3501(s string, queries [][]int) []int {
	n := len(s)
	cnt1 := 0
	for _, c := range s {
		if c == '1' {
			cnt1++
		}
	}

	var zeroBlocks []int
	var blockLeft []int
	var blockRight []int

	i := 0
	for i < n {
		st := i
		for i < n && s[i] == s[st] {
			i++
		}
		if s[st] == '0' {
			zeroBlocks = append(zeroBlocks, i-st)
			blockLeft = append(blockLeft, st)
			blockRight = append(blockRight, i-1)
		}
	}

	m := len(zeroBlocks)
	if m < 2 { // continuous 0 blocks less than 2 segments, return the answer directly
		ans := make([]int, len(queries))
		for i := range ans {
			ans[i] = cnt1
		}
		return ans
	}

	tmpSum := make([]int, m-1)
	for i := 0; i < m-1; i++ {
		tmpSum[i] = zeroBlocks[i] + zeroBlocks[i+1]
	}
	seg := NewSegmentTree3501(tmpSum)
	ans := make([]int, len(queries))

	for qIdx, q := range queries {
		l, r := q[0], q[1]
		i := sort.Search(len(blockRight), func(idx int) bool { return blockRight[idx] >= l })
		j := sort.Search(len(blockLeft), func(idx int) bool { return blockLeft[idx] > r }) - 1
		// at most 1 continuous block of 0s within the substring
		if i > m-1 || j < 0 || i >= j {
			ans[qIdx] = cnt1
			continue
		}

		firstLen := blockRight[i] - max(blockLeft[i], l) + 1 // actual length of the first consecutive block of 0s in the substring
		lastLen := min(blockRight[j], r) - blockLeft[j] + 1  // actual length of the last consecutive block of 0s in the substring
		// exactly 2 consecutive 0 blocks within the substring
		if i+1 == j {
			bestGain := firstLen + lastLen
			ans[qIdx] = cnt1 + bestGain
			continue
		}

		val1 := firstLen + zeroBlocks[i+1]
		val2 := zeroBlocks[j-1] + lastLen
		val3 := seg.Query(i+1, j-2)
		bestGain := max(val1, max(val2, val3))

		ans[qIdx] = cnt1 + bestGain
	}

	return ans
}
