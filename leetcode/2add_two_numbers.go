package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	nodeList := make([]*ListNode, 0)
	cf := 0
	for l1 != nil || l2 != nil {
		tmpNode := &ListNode{Val: 0}
		if l1 != nil {
			tmpNode.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmpNode.Val += l2.Val
			l2 = l2.Next
		}
		tmpNode.Val += cf
		cf = 0
		if tmpNode.Val >= 10 {
			cf = 1
			tmpNode.Val -= 10
		}
		nodeList = append(nodeList, tmpNode)
	}
	if cf != 0 {
		nodeList = append(nodeList, &ListNode{cf, nil})
	}

	if len(nodeList) == 0 {
		return nil
	}
	for idx := 0; idx < len(nodeList)-1; idx++ {
		nodeList[idx].Next = nodeList[idx+1]
	}
	return nodeList[0]
}
