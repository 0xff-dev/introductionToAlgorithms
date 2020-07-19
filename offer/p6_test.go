package offer

import "testing"

func TestConstruct(t *testing.T) {
	pre, in := []int{1,2,4,7,3,5,6,8}, []int{4,7,2,1,5,3,8,6}
	tree := Construct(pre, in)
	tree.PortOrder()
}
