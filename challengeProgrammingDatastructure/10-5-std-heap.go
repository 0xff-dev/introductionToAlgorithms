package challengeprogrammingdatastructure

import "container/heap"

type V []int

func (v *V) Len() int {
	return len(*v)
}

func (v *V) Less(i, j int) bool {
	return (*v)[i] > (*v)[j]
}

func (v *V) Swap(i, j int) {
	(*v)[i], (*v)[j] = (*v)[j], (*v)[i]
}

func (v *V) Push(x interface{}) {
	*v = append(*v, x.(int))
}

func (v *V) Pop() interface{} {
	old := *v
	l := len(*v)
	x := old[l-1]
	*v = old[:l-1]
	return x
}

func TryPriorityQuque(nums []int) []int {
	v := V([]int{})
	heap.Init(&v)
	for _, n := range nums {
		heap.Push(&v, n)
	}
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = heap.Pop(&v).(int)
	}
	return ans
}
