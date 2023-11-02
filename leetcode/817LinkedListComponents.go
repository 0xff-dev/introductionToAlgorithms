package leetcode

func numComponents(head *ListNode, nums []int) int {
	set := make(map[int]struct{})
	for _, n := range nums {
		set[n] = struct{}{}
	}
	ans := 0
	c := -1
	for walker := head; walker != nil; walker = walker.Next {
		if _, ok := set[walker.Val]; !ok {
			c = -1
			continue
		}
		if c == -1 {
			ans++
			c++
		}
	}
	return ans
}
