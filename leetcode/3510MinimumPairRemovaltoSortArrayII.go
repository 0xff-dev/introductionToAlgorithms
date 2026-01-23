package leetcode

import "container/heap"

func minimumPairRemoval3510(nums []int) int {
	pq := &PriorityQeueu3510{}
	heap.Init(pq)
	merged := make([]bool, len(nums))
	decreaseCount := 0
	count := 0
	head := &Node3510{value: int64(nums[0]), left: 0}
	current := head

	for i := 1; i < len(nums); i++ {
		newNode3510 := &Node3510{value: int64(nums[i]), left: i}
		current.next = newNode3510
		newNode3510.prev = current

		heap.Push(pq, &Item3510{
			first:  current,
			second: newNode3510,
			cost:   current.value + newNode3510.value,
		})
		if nums[i-1] > nums[i] {
			decreaseCount++
		}

		current = newNode3510
	}

	for decreaseCount > 0 {
		item := heap.Pop(pq).(*Item3510)
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

		prevNode3510 := first.prev
		nextNode3510 := second.next
		first.next = nextNode3510
		if nextNode3510 != nil {
			nextNode3510.prev = first
		}

		if prevNode3510 != nil {
			if prevNode3510.value > first.value && prevNode3510.value <= cost {
				decreaseCount--
			} else if prevNode3510.value <= first.value && prevNode3510.value > cost {
				decreaseCount++
			}
			heap.Push(pq, &Item3510{
				first:  prevNode3510,
				second: first,
				cost:   prevNode3510.value + cost,
			})
		}

		if nextNode3510 != nil {
			if second.value > nextNode3510.value && cost <= nextNode3510.value {
				decreaseCount--
			} else if second.value <= nextNode3510.value && cost > nextNode3510.value {
				decreaseCount++
			}
			heap.Push(pq, &Item3510{
				first:  first,
				second: nextNode3510,
				cost:   cost + nextNode3510.value,
			})
		}

		first.value = cost
		merged[second.left] = true
	}

	return count
}

type Node3510 struct {
	value int64
	left  int
	prev  *Node3510
	next  *Node3510
}

type Item3510 struct {
	first  *Node3510
	second *Node3510
	cost   int64
	index  int
}

type PriorityQeueu3510 []*Item3510

func (pq PriorityQeueu3510) Len() int { return len(pq) }
func (pq PriorityQeueu3510) Less(i, j int) bool {
	if pq[i].cost == pq[j].cost {
		return pq[i].first.left < pq[j].first.left
	}
	return pq[i].cost < pq[j].cost
}
func (pq PriorityQeueu3510) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQeueu3510) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item3510)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQeueu3510) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
