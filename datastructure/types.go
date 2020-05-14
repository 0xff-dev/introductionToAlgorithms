package datastructure

type Compare interface {
	Less(other interface{}) bool
	Swap(other interface{})
	Print()
}

type listNode struct {
	data Compare
	next *listNode
}

type List struct {
	Size int // list length
	Head *listNode
}
