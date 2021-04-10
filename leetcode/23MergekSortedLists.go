package leetcode

import "math"

// 2个一组合并，最后合并
func mergeKLists(lists []*ListNode) *ListNode {
	var newHead, list *ListNode
	length := len(lists)
	if length == 0 {
		return newHead
	}
	count := length
	for count > 0 {
		idx, min := -1, math.MaxInt32
		for index := 0; index < length; index++ {
			if lists[index] == nil {
				continue
			}
			if lists[index].Val < min {
				min = lists[index].Val
				idx = index
			}
		}
		if idx == -1 {
			return newHead
		}
		node := &ListNode{
			Val:  min,
			Next: nil,
		}
		if list == nil {
			newHead = node
			list = node
		} else {
			list.Next = node
			list = list.Next
		}
		lists[idx] = lists[idx].Next
		if lists[idx] == nil {
			count--
		}
	}

	return newHead
}
