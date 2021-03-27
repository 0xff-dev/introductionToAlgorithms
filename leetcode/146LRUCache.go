package leetcode

import (
	"container/list"
)

type LRUCache struct {
	cap   int
	l     *list.List
	store map[int]*list.Element
}

type elementVal struct {
	key, val int
}

func LRUConstructor(capacity int) LRUCache {
	return LRUCache{
		cap:   capacity,
		store: map[int]*list.Element{},
		l:     list.New(),
	}
}

func (this *LRUCache) get(key int) *list.Element {
	return this.store[key]
}

func (this *LRUCache) Get(key int) int {
	if val := this.get(key); val != nil {
		this.l.MoveToFront(val)
		return val.Value.(elementVal).val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	val := this.store[key]
	if val != nil {
		tmp := val.Value.(elementVal).val
		if tmp != value {
			val.Value = elementVal{key: key, val: value}
		}
		this.l.MoveToFront(val)
		return
	}
	if this.l.Len() >= this.cap {
		lastEle := this.l.Back()
		delete(this.store, (lastEle.Value).(elementVal).key)
		this.l.Remove(lastEle)
	}
	this.l.PushFront(elementVal{key: key, val: value})
	this.store[key] = this.l.Front()
}
