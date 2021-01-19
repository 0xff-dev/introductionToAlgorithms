package leetcode

import "testing"

func TestCopyRandomList(t *testing.T) {
	one := &Node{1, nil, nil}
	two := &Node{2, nil, nil}
	one.Next, one.Random = two, two
	two.Random = two
	head := copyRandomList(one)
	head.Dis()

	twoHead := copyRandomList(two)
	twoHead.Dis()
}
