package leetcode

import "sort"

type NumberContainers struct {
	// index对应的数字
	indies map[int]int
	// 所有的下标列表
	list map[int][]int
}

func Constructor2349() NumberContainers {
	return NumberContainers{
		indies: map[int]int{},
		list:   map[int][]int{},
	}
}

func (this *NumberContainers) Change(index int, number int) {
	v, ok := this.indies[index]
	this.indies[index] = number
	if !ok {
		// 新的数字
		list := this.list[number]
		idx := sort.Search(len(list), func(i int) bool {
			return list[i] > index
		})
		if idx == len(list) {
			this.list[number] = append(this.list[number], index)
		} else {
			// 1, 2, 3, [4] 5
			b := append([]int{index}, list[idx:]...)
			a := append(list[:idx], b...)
			this.list[number] = a
		}
		return
	}
	if v != number {
		list := this.list[v]
		idx := sort.Search(len(list), func(i int) bool {
			return list[i] >= index
		})
		this.list[v] = append(this.list[v][:idx], this.list[v][idx+1:]...)

		list = this.list[number]
		idx = sort.Search(len(list), func(i int) bool {
			return list[i] > index
		})
		if idx == len(list) {
			this.list[number] = append(this.list[number], index)
		} else {
			b := append([]int{index}, list[idx:]...)
			a := append(list[:idx], b...)
			this.list[number] = a
		}
	}

}

func (this *NumberContainers) Find(number int) int {
	v := this.list[number]
	if len(v) == 0 {
		return -1
	}
	return v[0]
}
