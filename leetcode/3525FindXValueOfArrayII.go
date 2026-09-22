package leetcode

func resultArray3525(nums []int, k int, queries [][]int) []int {
	n := len(nums)
	seg := NewSegmentTree35253525(nums, k)
	ans := []int{}

	for _, q := range queries {
		index, value, start, x := q[0], q[1], q[2], q[3]
		seg.Update(1, 0, n-1, index, value)
		pre := seg.Query(1, 0, n-1, start, n-1)
		ans = append(ans, pre[x])
	}
	return ans
}

type SegmentTree3525 struct {
	k    int
	tree [][]int
}

func NewSegmentTree35253525(nums []int, k int) *SegmentTree3525 {
	n := len(nums)
	size := 1 << (bitsLen(n) + 1)
	tree := make([][]int, size)
	for i := range tree {
		tree[i] = make([]int, k+1)
	}
	seg := &SegmentTree3525{k: k, tree: tree}
	seg.build(nums, 1, 0, n-1)
	return seg
}

func bitsLen(n int) int {
	cnt := 0
	for n > 0 {
		cnt++
		n >>= 1
	}
	return cnt
}

func (seg *SegmentTree3525) makeLeaf(o int, value int) {
	info := make([]int, seg.k+1)
	r := value % seg.k
	info[r] = 1
	info[seg.k] = r
	seg.tree[o] = info
}

func (seg *SegmentTree3525) mergePre(left, right []int) []int {
	pre := make([]int, seg.k+1)
	mulL := left[seg.k]
	mulR := right[seg.k]
	pre[seg.k] = (mulL * mulR) % seg.k

	for x := 0; x < seg.k; x++ {
		pre[x] = left[x]
	}
	for x := 0; x < seg.k; x++ {
		pre[(mulL*x)%seg.k] += right[x]
	}
	return pre
}

func (seg *SegmentTree3525) maintain(o int) {
	seg.tree[o] = seg.mergePre(seg.tree[o*2], seg.tree[o*2+1])
}

func (seg *SegmentTree3525) build(nums []int, o, l, r int) {
	if l == r {
		seg.makeLeaf(o, nums[l])
		return
	}
	m := (l + r) / 2
	seg.build(nums, o*2, l, m)
	seg.build(nums, o*2+1, m+1, r)
	seg.maintain(o)
}

func (seg *SegmentTree3525) Update(o, l, r, index, value int) {
	if l == r {
		seg.makeLeaf(o, value)
		return
	}
	m := (l + r) / 2
	if index <= m {
		seg.Update(o*2, l, m, index, value)
	} else {
		seg.Update(o*2+1, m+1, r, index, value)
	}
	seg.maintain(o)
}

func (seg *SegmentTree3525) Query(o, l, r, L, R int) []int {
	if L <= l && r <= R {
		return seg.tree[o]
	}
	m := (l + r) / 2
	if R <= m {
		return seg.Query(o*2, l, m, L, R)
	}
	if L > m {
		return seg.Query(o*2+1, m+1, r, L, R)
	}
	left := seg.Query(o*2, l, m, L, R)
	right := seg.Query(o*2+1, m+1, r, L, R)
	return seg.mergePre(left, right)
}
