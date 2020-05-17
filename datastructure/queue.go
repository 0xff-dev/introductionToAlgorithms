package datastructure

type queue struct {
	Size int
	Data []*treeNode
}

func NewQueue() *queue {
	return &queue{
		Size: 0,
		Data: []*treeNode{},
	}
}

func (self *queue) EnQueue(node *treeNode) {
	self.Data = append(self.Data, node)
	self.Size++
}

func (self *queue) Empty() bool {
	return self.Size == 0
}

func (self *queue) Peek() *treeNode {
	if !self.Empty() {
		return self.Data[0]
	}
	return nil
}

func (self *queue) DeQueue() *treeNode {
	if !self.Empty() {
		front := self.Data[0]
		self.Data = self.Data[1:]
		self.Size--
		return front
	}
	return nil
}
