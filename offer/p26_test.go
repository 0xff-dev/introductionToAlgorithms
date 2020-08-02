package offer

import "testing"

func TestComplexListReplication(t *testing.T) {

	A := &ComplexListNode{
		Val:   1,
		Next:  nil,
		Other: nil,
	}
	B := &ComplexListNode{
		Val:   2,
		Next:  nil,
		Other: nil,
	}
	C := &ComplexListNode{
		Val:   3,
		Next:  nil,
		Other: nil,
	}
	D := &ComplexListNode{
		Val:   4,
		Next:  nil,
		Other: nil,
	}
	E := &ComplexListNode{
		Val:   5,
		Next:  nil,
		Other: nil,
	}
	A.Next, A.Other = B, C
	B.Next, B.Other = C, E
	C.Next = D
	D.Next, D.Other = E, B
	A.Print()

	ComplexListReplication(A).Print()
}
