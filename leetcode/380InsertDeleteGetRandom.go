package leetcode

import "math/rand"

/*
	type RandomizedSet struct {
		l      *list.List
		bucket map[int]*list.Element
	}

// u32

	func Constructor380() RandomizedSet {
		return RandomizedSet{
			// 这里实际考虑的是直接用数组
			bucket: make(map[int]*list.Element),
			l:      list.New(),
		}
	}

	func (this *RandomizedSet) Insert(val int) bool {
		_, ok := this.bucket[val]
		if !ok {
			this.bucket[val] = this.l.PushBack(val)
		}
		return !ok
	}

	func (this *RandomizedSet) Remove(val int) bool {
		e, ok := this.bucket[val]
		if ok {
			this.l.Remove(e)
		}
		delete(this.bucket, val)
		return ok
	}

	func (this *RandomizedSet) GetRandom() int {
		// 遍历复杂度不是1啊
		for k := range this.bucket {
			return k
		}
		return 0
	}
*/
type RandomizedSet struct {
	bucket map[int]int
	stack  []int
	index  int
}

func Constructor380() RandomizedSet {
	return RandomizedSet{
		bucket: make(map[int]int),
		stack:  make([]int, 0),
		index:  -1,
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.bucket[val]
	if !ok {
		if this.index == len(this.stack)-1 {
			this.stack = append(this.stack, val)
			this.index++
		} else {
			this.index++
			this.stack[this.index] = val
		}
		this.bucket[val] = this.index
	}
	return !ok
}

func (this *RandomizedSet) Remove(val int) bool {
	removedIndex, ok := this.bucket[val]
	if ok {
		this.bucket[this.stack[this.index]] = removedIndex
		this.stack[removedIndex] = this.stack[this.index]

		this.index--
		delete(this.bucket, val)
	}
	return ok
}

func (this *RandomizedSet) GetRandom() int {
	i := rand.Intn(int(this.index) + 1)
	return this.stack[i]
}
