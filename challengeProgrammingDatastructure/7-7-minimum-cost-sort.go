package challengeprogrammingdatastructure

import "sort"

func MinimumCostSort(nums []int) int {
	b := make([]int, len(nums))
	v := make([]bool, len(nums))
	vmax := 0
	for i, n := range nums {
		if n > vmax {
			vmax = n
		}
		b[i] = n
	}
	sort.Ints(b)

	t := make([]int, vmax+1)
	for i := 0; i < len(nums); i++ {
		t[b[i]] = i
	}
	ans := 0

	for i := 0; i < len(nums); i++ {
		if v[i] {
			continue
		}
		cur := i
		s, m, an := 0, vmax, 0
		for {
			v[cur] = true
			an++
			vv := nums[cur]
			if m > vv {
				m = vv
			}
			s += vv
			cur = t[vv]
			if v[cur] {
				break
			}
		}
		a := s + (an-2)*m
		b := m + s + (an+1)*s
		if a > b {
			a = b
		}
		ans += a
	}
	return ans
}
