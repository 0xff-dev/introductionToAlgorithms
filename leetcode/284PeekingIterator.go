package leetcode

type Iterator struct{}

func (this *Iterator) hasNext() bool {
	return true
}
func (this *Iterator) next() int {
	return 0
}

type PeekingIterator struct {
	iter  *Iterator
	peekV int
}

func Constructor284(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter: iter, peekV: -1}
}

func (this *PeekingIterator) hasNext() bool {
	if !this.iter.hasNext() {
		return this.peekV != -1
	}
	return this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
	if this.peekV != -1 {
		ans := this.peekV
		this.peekV = -1
		return ans
	}
	return this.iter.next()
}

func (this *PeekingIterator) peek() int {
	if this.peekV == -1 {
		this.peekV = this.iter.next()
	}
	return this.peekV
}
