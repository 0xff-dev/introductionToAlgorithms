package leetcode

import (
	"github.com/emirpasic/gods/trees/redblacktree"
)

type MultiSet struct {
	tree    *redblacktree.Tree
	counter map[int]int
	size    int
}

func NewMultiSet() *MultiSet {
	return &MultiSet{
		tree:    redblacktree.NewWithIntComparator(),
		counter: make(map[int]int),
		size:    0,
	}
}

func (ms *MultiSet) Add(x int) {
	if count, exists := ms.counter[x]; exists {
		ms.counter[x] = count + 1
	} else {
		ms.counter[x] = 1
		ms.tree.Put(x, struct{}{})
	}
	ms.size++
}

func (ms *MultiSet) Remove(x int) bool {
	if count, exists := ms.counter[x]; exists {
		if count == 1 {
			delete(ms.counter, x)
			ms.tree.Remove(x)
		} else {
			ms.counter[x] = count - 1
		}
		ms.size--
		return true
	}
	return false
}

func (ms *MultiSet) Size() int {
	return ms.size
}

func (ms *MultiSet) IsEmpty() bool {
	return ms.size == 0
}

func (ms *MultiSet) First() (int, bool) {
	if ms.tree.Empty() {
		return 0, false
	}
	return ms.tree.Left().Key.(int), true
}

func (ms *MultiSet) Last() (int, bool) {
	if ms.tree.Empty() {
		return 0, false
	}
	return ms.tree.Right().Key.(int), true
}

func (ms *MultiSet) Contains(x int) bool {
	_, exists := ms.counter[x]
	return exists
}

type Container struct {
	k   int
	st1 *MultiSet
	st2 *MultiSet
	sm  int64
}

func NewContainer(k int) *Container {
	return &Container{
		k:   k,
		st1: NewMultiSet(),
		st2: NewMultiSet(),
		sm:  0,
	}
}

func (m *Container) adjust() {
	for m.st1.Size() < m.k && !m.st2.IsEmpty() {
		if x, ok := m.st2.First(); ok {
			m.st2.Remove(x)
			m.st1.Add(x)
			m.sm += int64(x)
		}
	}
	for m.st1.Size() > m.k {
		if x, ok := m.st1.Last(); ok {
			m.st1.Remove(x)
			m.st2.Add(x)
			m.sm -= int64(x)
		}
	}
}

// insert element x
func (m *Container) add(x int) {
	if !m.st2.IsEmpty() {
		if first, ok := m.st2.First(); ok && x >= first {
			m.st2.Add(x)
		} else {
			m.st1.Add(x)
			m.sm += int64(x)
		}
	} else {
		m.st1.Add(x)
		m.sm += int64(x)
	}
	m.adjust()
}

// delete element x
func (m *Container) erase(x int) {
	if m.st1.Contains(x) {
		m.st1.Remove(x)
		m.sm -= int64(x)
	} else if m.st2.Contains(x) {
		m.st2.Remove(x)
	}
	m.adjust()
}

// sum of the first k smallest elements
func (m *Container) sum() int64 {
	return m.sm
}

func minimumCost3013(nums []int, k int, dist int) int64 {
	n := len(nums)
	cnt := NewContainer(k - 2)

	for i := 1; i < k-1; i++ {
		cnt.add(nums[i])
	}

	ans := cnt.sum() + int64(nums[k-1])
	for i := k; i < n; i++ {
		j := i - dist - 1
		if j > 0 {
			cnt.erase(nums[j])
		}
		cnt.add(nums[i-1])
		current := cnt.sum() + int64(nums[i])
		if current < ans {
			ans = current
		}
	}

	return ans + int64(nums[0])
}
