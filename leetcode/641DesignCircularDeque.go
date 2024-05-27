package leetcode

import "container/list"

type MyCircularDeque struct {
	k    int
	list *list.List
}

func Constructor641(k int) MyCircularDeque {
	return MyCircularDeque{
		k:    k,
		list: list.New(),
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.list.Len() == this.k {
		return false
	}
	this.list.PushFront(value)
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.list.Len() == this.k {
		return false
	}
	this.list.PushBack(value)
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.list.Len() == 0 {
		return false
	}
	head := this.list.Front()
	this.list.Remove(head)
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.list.Len() == 0 {
		return false
	}
	head := this.list.Back()
	this.list.Remove(head)
	return true

}

func (this *MyCircularDeque) GetFront() int {
	if this.list.Len() == 0 {
		return -1
	}
	return this.list.Front().Value.(int)
}

func (this *MyCircularDeque) GetRear() int {
	if this.list.Len() == 0 {
		return -1
	}
	return this.list.Back().Value.(int)

}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.list.Len() == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.k == this.list.Len()
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
