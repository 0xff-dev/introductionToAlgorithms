package leetcode

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return head
	}
	idx, walker, last := 1, head, n-m
	numArray := make([]int, n-m+1)
	for ; walker != nil && idx <= n; walker, idx = walker.Next, idx+1 {
		if idx >= m {
			numArray[last] = walker.Val
			last--
		}
	}
	index := 0
	for walker, idx := head, 1; walker != nil; walker, idx = walker.Next, idx+1 {
		if idx > n {
			break
		}
		if idx >= m {
			walker.Val = numArray[index]
			index++
		}
	}
	return head
}
