package leetcode

func modifiedList(nums []int, head *ListNode) *ListNode {
	in := make(map[int]struct{})
	for _, n := range nums {
		in[n] = struct{}{}
	}
	walker := head
	var preWalker *ListNode
	for walker != nil {
		if _, ok := in[walker.Val]; !ok {
			preWalker = walker
			walker = walker.Next
			continue
		}
		if walker.Next == nil {
			preWalker.Next = nil
			break
		}
		*walker = *walker.Next
	}
	return head
}
