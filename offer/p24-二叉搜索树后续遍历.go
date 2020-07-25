package offer

import "fmt"

func PostOrderRecursive(tree *TreeNode) {
	if tree.Left != nil {
		PostOrderRecursive(tree.Left)
	}
	if tree.Right != nil {
		PostOrderRecursive(tree.Right)
	}
	fmt.Println("data: ", tree.Val)
}

type treeStack struct {
	Data []*TreeNode
}

func (ts *treeStack) Push(node *TreeNode) {
	ts.Data = append(ts.Data, node)
}
func (ts *treeStack) Top() (*TreeNode, error) {
	if len(ts.Data) == 0 {
		return nil, fmt.Errorf("empty stack")
	}
	return ts.Data[len(ts.Data)-1], nil
}
func (ts *treeStack) Pop() (*TreeNode, error) {
	top, err := ts.Top()
	if err != nil {
		return nil, err
	}
	ts.Data = ts.Data[:len(ts.Data)-1]
	return top, nil
}

func PostOrderStack(tree *TreeNode) {
	s := &treeStack{}
	for tree != nil || len(s.Data) != 0 {
		for ; tree != nil; tree = tree.Left {
			s.Push(tree)
			s.Push(tree)
		}
		if len(s.Data) != 0 {
			_top, _ := s.Pop()
			if top, err := s.Top(); err == nil && top == _top {
				tree = _top.Right
			} else {
				fmt.Println("data: ", _top.Val)
				tree = nil
			}
		}
	}
}
