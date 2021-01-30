package leetcode

import "fmt"

type MyLinkedList struct {
	Data []int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{Data: []int{}}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index >= len(this.Data) {
		return -1
	}
	return this.Data[index]
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	this.Data = append([]int{val}, this.Data...)
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	this.Data = append(this.Data, val)
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > len(this.Data) {
		return
	}

	this.Data = append(this.Data[:index], append([]int{val}, this.Data[index:]...)...)
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index > len(this.Data) {
		return
	}
	fmt.Println("delete index: ", index)
	if index+1 >= len(this.Data) {
		this.Data = this.Data[:index]
		return
	}
	this.Data = append(this.Data[:index], this.Data[index+1:]...)
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
