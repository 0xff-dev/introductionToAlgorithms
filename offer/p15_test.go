package offer

import (
	"fmt"
	"testing"
)

func TestKthToTail(t *testing.T) {
	list := &ListNode{1, &ListNode{2, nil}}
	result := KthToTail(list, 3)
	fmt.Println(result)
}
