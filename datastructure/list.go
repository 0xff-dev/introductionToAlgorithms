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
