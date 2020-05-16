package datastructure

import "fmt"

// this stack only for tree, we can change it to others types
type stack struct {
	Size int
	Data []*treeNode
}

func (self *stack) Empty() bool {
	return self.Size == 0
}

func (self *stack) Push(data *treeNode) {
	self.Data = append(self.Data, data)
	self.Size++
}

func (self *stack) Top() (*treeNode, error) {
	if !self.Empty() {
		return self.Data[self.Size-1], nil
	}
	return nil, fmt.Errorf("empty stack")
}

func (self *stack) Pop() (*treeNode, error) {
	if !self.Empty() {
		self.Size--
		topData := self.Data[self.Size]
		self.Data = self.Data[:self.Size]
		return topData, nil
	}
	return nil, fmt.Errorf("empty stack")
}

func NewStack() *stack {
	return &stack{
		Size: 0,
		Data: []*treeNode{},
	}
}
