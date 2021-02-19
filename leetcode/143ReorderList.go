package leetcode

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	nodeArray := make([]*ListNode, 0)
	for walker := head; walker != nil; walker = walker.Next {
		nodeArray = append(nodeArray, walker)
	}
	start, end := 0, len(nodeArray)-1
	for start < end {
		nodeArray[start].Next = nodeArray[end]
		start, end = start+1, end-1
		nodeArray[start-1].Next.Next = nodeArray[start]
		if start == end || start == end-1 {
			nodeArray[end].Next = nil
			break
		}
	}

}
