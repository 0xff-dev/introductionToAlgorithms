package leetcode

import (
	"math/rand"
	"time"
)

type Solution382 struct {
	length int
	head   *ListNode
}

func Constructor382(head *ListNode) Solution382 {
	s := Solution382{length: 0, head: head}
	for walker := head; walker != nil; walker = walker.Next {
		s.length++
	}
	return s
}

func (this *Solution382) GetRandom() int {
	if this.head == nil {
		return -1
	}
	seed := rand.NewSource(time.Now().UnixNano())
	rr := rand.New(seed)
	index := rr.Intn(this.length)

	walker := this.head
	for ; walker != nil && index > 0; walker = walker.Next {
		index--
	}
	return walker.Val
}
