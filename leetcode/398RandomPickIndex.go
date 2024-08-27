package leetcode

type item398 struct {
	index int
	list  []int
}

type Solution398 struct {
	data map[int]item398
}

func Constructor398(nums []int) Solution398 {
	s := Solution398{data: make(map[int]item398)}
	for i, n := range nums {
		if _, ok := s.data[n]; !ok {
			s.data[n] = item398{}
		}
		l := s.data[n]
		l.list = append(l.list, i)
		s.data[n] = l
	}
	return s
}

func (this *Solution398) Pick(target int) int {
	v := this.data[target]
	cur := v.index
	v.index = (v.index + 1) % len(v.list)
	this.data[target] = v
	return this.data[target].list[cur]
}
