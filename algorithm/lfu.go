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
	header, tail *lfuNode
	size         int // page size
	length       int // length of list
}

func NewLFUList() *lfuList {
	return &lfuList{header: nil, size: 5, length: 0}
}

func (l *lfuList) Add(node *lfuNode) {
	if l.header == nil {
		l.header = node
		l.tail = l.header
		return
	}

	if l.length == l.size {
		// todo delete tail node
		l.Delete(l.tail)
	}
	walker := l.header
	for ; walker.Next != nil; walker = walker.Next {
	}
	walker.Next = node
	l.tail = node
	// add to header
	l.length++
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

func (l *lfuList) Find(key int) *lfuNode {
	walker := l.header
	for ; walker != nil; walker = walker.Next {
		if walker.Val == key {
			return walker
		}
	}
	return nil
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

func (lfu *lfuList) Get(key int) *lfuNode {
	node := lfu.Find(key)
	if node == nil {
		return nil
	}

	// add fre
	node.Fre++
	lfu.Sort()
	return node
}

func (lfu *lfuList) Set(key, newKey int) {
	node := lfu.Find(key)
	if node == nil {
		n := &lfuNode{
			Fre:     1,
			Val:     newKey,
			AddTime: time.Now(),
			Next:    nil,
		}
		lfu.Add(n)
		return
	}
	node.Fre++
	node.Val = newKey
	lfu.Sort()
}
