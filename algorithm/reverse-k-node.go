package algorith

import "fmt"

/*
这其实是一道变形的链表反转题，大致描述如下 给定一个单链表的头节点 head,实现一个调整单链表的函数，
使得每K个节点之间为一组进行逆序，并且从链表的尾部开始组起，头部剩余节点数量不够一组的不需要逆序。（
不能使用队列或者栈作为辅助）
*/

type listNode struct {
	Val  int
	Next *listNode
}

func (list *listNode) Reverse() *listNode {
	if list == nil || list.Next == nil {
		return list
	}
	newNode := list.Next.Reverse()
	list.Next.Next = list
	list.Next = nil
	return newNode
}

func (list *listNode) Dis() {
	for p := list; p != nil; p = p.Next {
		fmt.Printf("val: %d", p.Val)
	}
	fmt.Println()
}

func ReverseKNode(head *listNode, k int) *listNode {
	if head == nil || k == 0 {
		return head
	}
	head = head.Reverse()
	walker, ps, cnt := head, head, 1
	var hs, pre *listNode
	for walker != nil {
		if cnt == k {
			if walker != nil {
				pre = walker
				walker = walker.Next
			}
			pre.Next = hs
			hs, ps = ps, walker
			cnt = 1
			continue
		}
		pre = walker
		walker = walker.Next
		cnt++
	}
	if cnt != 1 {
		tmpHead := ps.Reverse()
		ps.Next = hs
		hs = tmpHead
	}
	return hs
}
