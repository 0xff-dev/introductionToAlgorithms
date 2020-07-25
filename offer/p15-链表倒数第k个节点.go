package offer

// 注意代码鲁棒性，空指针问题
func KthToTail(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return nil
	}
	former := head
	for index := 0; index < k-1 && former != nil; index, former = index+1, former.Next {
	}
	if former == nil {
		return nil
	}
	latter := head
	for ; former.Next != nil; former, latter = former.Next, latter.Next {
	}
	return latter
}
