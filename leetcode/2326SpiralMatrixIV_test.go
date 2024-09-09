package leetcode

import (
	"reflect"
	"testing"
)

func TestSpiralMatrix(t *testing.T) {
	m, n, head := 3, 5, &ListNode{3, &ListNode{0, &ListNode{2, &ListNode{6, &ListNode{8, &ListNode{1, &ListNode{7, &ListNode{9, &ListNode{4, &ListNode{2, &ListNode{5, &ListNode{5, &ListNode{0, nil}}}}}}}}}}}}}
	exp := [][]int{
		{3, 0, 2, 6, 8},
		{5, 0, -1, -1, 1},
		{5, 2, 4, 9, 7},
	}
	if r := spiralMatrix(m, n, head); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	m, n, head = 1, 4, &ListNode{0, &ListNode{1, &ListNode{2, nil}}}
	exp = [][]int{
		{0, 1, 2, -1},
	}
	if r := spiralMatrix(m, n, head); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
