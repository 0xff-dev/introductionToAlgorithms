package leetcode

type FenwickTree struct {
	tree []int
}

func fenwickTree(size int) *FenwickTree {
	return &FenwickTree{tree: make([]int, size+1)}
}

func (ft *FenwickTree) update(index, delta int) {
	index++
	for index < len(ft.tree) {
		ft.tree[index] += delta
		index += index & -index
	}
}

func (ft *FenwickTree) query(index int) int {
	index++
	res := 0
	for index > 0 {
		res += ft.tree[index]
		index -= index & -index
	}
	return res
}

func goodTriplets(nums1 []int, nums2 []int) int64 {
	n := len(nums1)
	pos2, reversedIndexMapping := make([]int, n), make([]int, n)
	for i, num := range nums2 {
		pos2[num] = i
	}

	for i, num := range nums1 {
		reversedIndexMapping[pos2[num]] = i
	}
	tree := fenwickTree(n)
	var res int64
	for value := 0; value < n; value++ {
		pos := reversedIndexMapping[value]
		left := tree.query(pos)
		tree.update(pos, 1)
		right := (n - 1 - pos) - (value - left)
		res += int64(left * right)
	}
	return res
}
