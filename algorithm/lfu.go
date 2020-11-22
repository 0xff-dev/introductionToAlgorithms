package algorith

import (
	"fmt"
	"time"
)

// 一个单连标+sort排序，
type lfuNode struct {
	Fre     int
	Val     int
	AddTime time.Time
	Next    *lfuNode
}

func (ln *lfuNode) Less(other *lfuNode) bool {
	if ln.Fre == other.Fre {
		return ln.AddTime.Before(other.AddTime)
	}
	return ln.Fre < other.Fre
}

func (ln *lfuNode) Swap(other *lfuNode) {
	ln.Val, other.Val = other.Val, ln.Val
	ln.Fre, other.Fre = other.Fre, ln.Fre
	ln.AddTime, other.AddTime = other.AddTime, ln.AddTime
}

type lfuList struct {
	header *lfuNode
}

func NewLFUList() *lfuList {
	return &lfuList{header: nil}
}

func (l *lfuList) Add(node *lfuNode) {
	if l.header == nil {
		l.header = node
		return
	}
	// add to header
	node.Next = l.header
	l.header = node

	l.Sort()
}

func (l *lfuList) Delete(node *lfuNode) {
	if node == l.header {
		l.header = l.header.Next
		return
	}
	walker := &l.header
	for *walker != nil {
		entry := *walker
		if entry == node {
			*walker = entry.Next
			return
		}
		walker = &(entry.Next)
	}
}

func (l *lfuList) Dis() {
	walker := l.header
	for walker != nil {
		fmt.Printf("Fre: %d, Val: %d, Time: %s\n", walker.Fre, walker.Val, walker.AddTime)
		walker = walker.Next
	}
}

// sort object by frequency
func (lfu *lfuList) Sort() {
	sort(lfu.header, nil)
}

func sort(start, end *lfuNode) {
	if start == end || start.Next == end {
		return
	}

	aim := start
	pre, walker := start, start.Next
	for walker != nil {
		if walker.Less(aim) {
			pre = pre.Next
			walker.Swap(pre)
		}
		walker = walker.Next
	}
}

// todo add lfuCache struct and impl get, set
