package leetcode

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	extendList(head)
	//head.Dis()
	addRandomPointer(head)
	_, head = splitList(head)
	return head
}

func extendList(head *Node) {
	for walker := head; walker != nil; {
		tmpNode := &Node{
			Val:    walker.Val,
			Next:   walker.Next,
			Random: nil,
		}
		walker.Next = tmpNode
		walker = tmpNode.Next
	}
}

func addRandomPointer(head *Node) {
	walker := head
	for walker != nil {
		if walker.Random != nil {
			walkerRandomNode := walker.Random
			nextSameNode := walker.Next
			nextSameNode.Random = walkerRandomNode.Next
		}
		walker = walker.Next.Next
	}
}

func splitList(head *Node) (*Node, *Node) {
	oldHeader, newHeader := head, head.Next
	w1, w2 := head, head.Next
	for w1 != nil {
		w1.Next = w2.Next
		w1 = w1.Next
		if w1 != nil {
			w2.Next = w1.Next
			w2 = w2.Next
			continue
		}
		w2.Next = nil
	}
	return oldHeader, newHeader
}
