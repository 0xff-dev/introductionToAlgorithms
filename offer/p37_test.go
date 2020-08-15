package offer

import "testing"

func TestCommonNode(t *testing.T) {
	commList := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	}
	list1 := &ListNode{
		Val: 8,
		Next: &ListNode{
			Val:  10,
			Next: commList,
		},
	}
	commonNode := CommonNode(list1, commList)
	if commonNode == nil || commonNode.Val != 3 {
		t.Fatalf("not common node")
	}
	t.Logf("successfully found common node: %d", commonNode.Val)
	commonNode = CommonNode(commonNode, commonNode)
	if commonNode == nil || commonNode.Val != 3 {
		t.Fatalf("expect commond node val: %d", 3)
	}
	t.Logf("successfully found common node: %d", commonNode.Val)
}
