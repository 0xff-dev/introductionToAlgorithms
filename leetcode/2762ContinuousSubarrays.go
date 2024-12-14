package leetcode

import "sort"

type SegmentTreeNode2762 struct {
	Left, Right int
	Min, Max    int
	LeftChild   *SegmentTreeNode2762
	RightChild  *SegmentTreeNode2762
}

func buildSegmentTree2762(arr []int, left, right int) *SegmentTreeNode2762 {
	if left == right {
		return &SegmentTreeNode2762{
			Left:  left,
			Right: right,
			Min:   arr[left],
			Max:   arr[left],
		}
	}

	mid := (left + right) / 2
	leftChild := buildSegmentTree2762(arr, left, mid)
	rightChild := buildSegmentTree2762(arr, mid+1, right)

	node := &SegmentTreeNode2762{
		Left:       left,
		Right:      right,
		Min:        min(leftChild.Min, rightChild.Min),
		Max:        max(leftChild.Max, rightChild.Max),
		LeftChild:  leftChild,
		RightChild: rightChild,
	}

	return node
}

func queryMin2762(node *SegmentTreeNode2762, left, right int) int {
	if node.Left == left && node.Right == right {
		return node.Min
	}

	mid := (node.Left + node.Right) / 2

	if right <= mid {
		return queryMin2762(node.LeftChild, left, right)
	} else if left > mid {
		return queryMin2762(node.RightChild, left, right)
	} else {
		leftMin := queryMin2762(node.LeftChild, left, mid)
		rightMin := queryMin2762(node.RightChild, mid+1, right)
		return min(leftMin, rightMin)
	}
}

func queryMax2762(node *SegmentTreeNode2762, left, right int) int {
	if node.Left == left && node.Right == right {
		return node.Max
	}

	mid := (node.Left + node.Right) / 2

	if right <= mid {
		return queryMax2762(node.LeftChild, left, right)
	} else if left > mid {
		return queryMax2762(node.RightChild, left, right)
	} else {
		leftMax := queryMax2762(node.LeftChild, left, mid)
		rightMax := queryMax2762(node.RightChild, mid+1, right)
		return max(leftMax, rightMax)
	}
}

func continuousSubarrays(nums []int) int64 {
	l := len(nums)
	ans := int64(1)

	tree := buildSegmentTree2762(nums, 0, l-1)
	// 这个算法可以通过, 感觉这个东西可以通过滑动窗口进行优化，等我回来的, 我先去溜达
	for end := 1; end < l; end++ {
		i := sort.Search(end, func(i int) bool {
			_max := queryMax2762(tree, i, end)
			_min := queryMin2762(tree, i, end)
			return _max-_min <= 2
		})
		ans += int64(end-i) + 1
	}

	return ans
}
