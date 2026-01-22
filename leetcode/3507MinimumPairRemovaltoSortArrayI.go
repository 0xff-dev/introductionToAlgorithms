package leetcode

import "container/heap"

func minimumPairRemoval(nums []int) int {
	pq := &PriorityQueue3507{}
	heap.Init(pq)
	merged := make([]bool, len(nums))
	decreaseCount := 0
	count := 0
	head := &Node3507{value: int64(nums[0]), left: 0}
	current := head

	for i := 1; i < len(nums); i++ {
		newNode3507 := &Node3507{value: int64(nums[i]), left: i}
		current.next = newNode3507
		newNode3507.prev = current

		heap.Push(pq, &Item{
			first:  current,
			second: newNode3507,
			cost:   current.value + newNode3507.value,
		})
		if nums[i-1] > nums[i] {
			decreaseCount++
		}

		current = newNode3507
	}

	for decreaseCount > 0 {
		item := heap.Pop(pq).(*Item)
		first := item.first
		second := item.second
		cost := item.cost

		if merged[first.left] || merged[second.left] || first.value+second.value != cost {
			continue
		}
		count++
		if first.value > second.value {
			decreaseCount--
		}

		prevNode3507 := first.prev
		nextNode3507 := second.next
		first.next = nextNode3507
		if nextNode3507 != nil {
			nextNode3507.prev = first
		}

		if prevNode3507 != nil {
			if prevNode3507.value > first.value && prevNode3507.value <= cost {
				decreaseCount--
			} else if prevNode3507.value <= first.value && prevNode3507.value > cost {
				decreaseCount++
			}
			heap.Push(pq, &Item{
				first:  prevNode3507,
				second: first,
				cost:   prevNode3507.value + cost,
			})
		}

		if nextNode3507 != nil {
			if second.value > nextNode3507.value && cost <= nextNode3507.value {
				decreaseCount--
			} else if second.value <= nextNode3507.value && cost > nextNode3507.value {
				decreaseCount++
			}
			heap.Push(pq, &Item{
				first:  first,
				second: nextNode3507,
				cost:   cost + nextNode3507.value,
			})
		}

		first.value = cost
		merged[second.left] = true
	}

	return count
}

type Node3507 struct {
	value int64
	left  int
	prev  *Node3507
	next  *Node3507
}

type Item struct {
	first  *Node3507
	second *Node3507
	cost   int64
	index  int
}

type PriorityQueue3507 []*Item

func (pq PriorityQueue3507) Len() int { return len(pq) }
func (pq PriorityQueue3507) Less(i, j int) bool {
	if pq[i].cost == pq[j].cost {
		return pq[i].first.left < pq[j].first.left
	}
	return pq[i].cost < pq[j].cost
}
func (pq PriorityQueue3507) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue3507) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue3507) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
