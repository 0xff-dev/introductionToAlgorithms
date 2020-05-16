package datastructure

func (self *listNode) GetData() interface{} {
	return self.data
}

func New() *List {
	return &List{
		Size: 0,
		Head: nil,
	}
}

func (self *List) InsertTail(data Compare) {
	newNode := &listNode{
		data: data,
		next: nil,
	}
	self.Size++
	if self.Head == nil {
		self.Head = newNode
		return
	}
	pre, walker := self.Head, self.Head.next
	for ; walker != nil; pre, walker = pre.next, walker.next {
	}
	pre.next = newNode
}

func (self *List) Sort() {
	if self.Head == nil {
		return
	}
	self.sort(self.Head, nil)
}

func (self *List) sort(start, end *listNode) {
	if start == end || start.next == end {
		return
	}
	compareNode := start
	pre, walker := start, start.next
	for walker != nil {
		if walker.data.Less(compareNode.data) {
			pre = pre.next
			walker.data.Swap(pre.data)
		}
		walker = walker.next
	}
	pre.data.Swap(compareNode.data)
	self.sort(start, pre)
	self.sort(pre.next, end)
}

func (self *List) Print() {
	for walker := self.Head; walker != nil; walker = walker.next {
		walker.data.Print()
	}
}

func (self *List) SetCircle(index int) {
	step, walker, aimNode := 0, self.Head, self.Head
	for ; walker.next != nil; walker, step = walker.next, step+1 {
		if step == index {
			aimNode = walker
		}
	}
	walker.next = aimNode
}

func (self *List) HasCircle() (bool, *listNode) {
	if self.Head == nil {
		return false, nil
	}
	hasCycle := false
	walker, nextWalker := self.Head, self.Head
	for nextWalker != nil && nextWalker.next != nil {
		nextWalker = nextWalker.next.next
		walker = walker.next
		if nextWalker == walker {
			hasCycle = true
			break
		}
	}
	return hasCycle, walker
}

func (self *List) CircleNode() *listNode {
	hasCycle, node := self.HasCircle()
	if hasCycle {
		walker := self.Head
		for walker != node {
			walker, node = walker.next, node.next
		}
		return node
	}
	return nil
}

func (self *List) Reverse(recursive bool) {
	if recursive {
		self.Head = self.Head.reverse()
		return
	}
	if self.Head == nil || self.Head.next == nil {
		return
	}
	pre, next := self.Head, self.Head.next
	pre.next = nil
	for next != nil {
		tmp := next.next
		next.next = pre
		pre, next = next, tmp
	}
	self.Head = pre
}
