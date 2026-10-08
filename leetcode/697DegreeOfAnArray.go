package leetcode

type elem struct {
	s, e, c int
}

func findShortestSubArray(nums []int) int {
	group := make(map[int]elem)
	degree := 0
	for i := range nums {
		v, ok := group[nums[i]]
		if !ok {
			v = elem{s: i, e: i, c: 1}
		}
		v.e = i
		v.c++
		group[nums[i]] = v
	}
	for _, v := range group {
		degree = max(degree, v.c)
	}
	ret := 1000000007
	for _, v := range group {
		if v.c != degree {
			continue
		}
		ret = min(ret, v.e-v.s+1)
	}
	return ret
}
