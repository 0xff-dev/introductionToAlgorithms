package leetcode

func countSubarrays2962(nums []int, k int) int64 {
	ans := int64(0)
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > m {
			m = nums[i]
		}
	}
	indies := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] == m {
			indies = append(indies, i)
		}
	}
	ll := len(nums)
	if ll < k {
		return ans
	}
	le := 0
	for s, e := 0, k-1; e < len(indies); s, e = s+1, e+1 {
		left := int64(indies[s] - le)
		right := int64(ll - 1 - indies[e])
		ans += left + right + left*right + 1
		le = indies[s] + 1
	}
	return ans
}

/*
type segNode2962 struct {
	l, r, m, n  int
	left, right *segNode2962
}

func (s *segNode2962) Floor() {
	q := []*segNode2962{s}
	for len(q) > 0 {
		next := make([]*segNode2962, 0)
		for _, i := range q {
			fmt.Printf("[%d, %d], max: %d ***$", i.l, i.r, i.m)
			if i.left != nil {
				next = append(next, i.left)
			}
			if i.right != nil {
				next = append(next, i.right)
			}
		}
		q = next
	}
}
func buildSegTree2962(left, right int, nums []int) *segNode2962 {
	node := &segNode2962{l: left, r: right}
	if left == right {
		node.m = nums[left]
		node.n = nums[left]
		return node
	}
	mid := (right + left) / 2
	node.left = buildSegTree2962(left, mid, nums)
	node.right = buildSegTree2962(mid+1, right, nums)
	node.m = max(node.left.m, node.right.m)
	node.n = min(node.left.n, node.right.n)
	return node
}

func maxSeg2962(root *segNode2962, s, e int, nums []int) int {
	if root.l >= s && root.r <= e {
		return root.m
	}
	if root.l > e || root.r < s {
		return -1
	}
	return max(maxSeg2962(root.left, s, e, nums), maxSeg2962(root.right, s, e, nums))
}
func minSeg2962(root *segNode2962, s, e int, nums []int) int {
	if root.l >= s && root.r <= e {
		return root.n
	}
	if root.l > e || root.r < s {
		return 1000001
	}
	return min(minSeg2962(root.left, s, e, nums), minSeg2962(root.right, s, e, nums))

}

// 感觉可以结合线段树+二分搜索
// 当我们确定了最少的边界后，就可以
func countSubarrays2962(nums []int, k int) int64 {
	ans := int64(0)
	// 我以nums[i] 作为最大值
	// 但是往回找，那就是n^2算法了，不一定能过
	//	root := buildSegTree2962(0, len(nums)-1, nums)

	indies := make(map[int][]int)
	for i, n := range nums {
		if _, ok := indies[n]; !ok {
			indies[n] = make([]int, 0)
		}
		indies[n] = append(indies[n], i)
	}
	tree := buildSegTree2962(0, len(nums)-1, nums)
	for mi, loc := range indies {
		fmt.Printf("--- mi: %d --- %v\n", mi, loc)
		l := len(loc)
		if l < k {
			continue
		}
		le, re := 0, len(nums)-1
		s, e := 0, k-1
		for ; e < l; s, e = s+1, e+1 {
			li := loc[s]
			ri := loc[e]

			target := maxSeg2962(tree, li, ri, nums)
			fmt.Printf("--- [%d - %d] max = %d\n", li, ri, target)
			if target != mi {
				le = loc[s] + 1
				continue
			}
			// 在这区间，最大值是mi,然后
			// 0, 1, 2, 3, 4,

			ll := li - le + 1
			fmt.Printf("-- left length = %d\n", ll)
			leftIndex := sort.Search(ll, func(i int) bool {
				return maxSeg2962(tree, le+i, li, nums) <= mi
			})

			left := ll - leftIndex - 1
			fmt.Printf("ll=%d, leftIndex = %d left= %d\n", ll, le+leftIndex, left)

			rl := re - ri + 1
			rightIndex := sort.Search(rl, func(i int) bool {
				return minSeg2962(tree, ri+i, re, nums) > mi
			})
			right := rightIndex - 1
			fmt.Printf("rl=%d, rightIndex=%d, right=%d\n", rl, rightIndex, right)
			cur := int64(left) + int64(right) + int64(left*right) + 1
			fmt.Printf("coutn = %d\n", cur)
			ans += cur
			//ans++
			le = loc[s] + 1
			fmt.Println()
		}
	}
	return ans
}
*/
