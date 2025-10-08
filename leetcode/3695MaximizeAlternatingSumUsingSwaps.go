package leetcode

import "sort"

type unionFind3695 struct {
	father map[int]int
}

func (u *unionFind3695) find(x int) int {
	if u.father[x] != x {
		u.father[x] = u.find(u.father[x])
	}
	return u.father[x]
}

func (u *unionFind3695) union(x, y int) {
	fx := u.find(x)
	fy := u.find(y)
	if fx < fy {
		u.father[fy] = fx
	} else {
		u.father[fx] = fy
	}
}

// 1, 1, 2
func maxAlternatingSum(nums []int, swaps [][]int) int64 {
	u := unionFind3695{
		father: make(map[int]int),
	}
	for _, swap := range swaps {
		u.father[swap[0]] = swap[0]
		u.father[swap[1]] = swap[1]
	}
	for _, swap := range swaps {
		a, b := swap[0], swap[1]
		u.union(a, b)
	}
	// 已经组合完了。开干
	child := make(map[int]map[int]struct{})
	for _, swap := range swaps {
		a, b := swap[0], swap[1]
		fa := u.find(a)
		if _, ok := child[fa]; !ok {
			child[fa] = make(map[int]struct{})
		}
		child[fa][a] = struct{}{}
		child[fa][b] = struct{}{}
	}
	for _, indies := range child {
		pickup := make([]int, len(indies))
		i := 0
		for index := range indies {
			pickup[i] = nums[index]
			i++
		}
		sort.Slice(pickup, func(i, j int) bool {
			return pickup[i] > pickup[j]
		})
		start, end := 0, len(indies)-1
		for index := range indies {
			if index&1 == 0 {
				nums[index] = pickup[start]
				start++
				continue
			}
			nums[index] = pickup[end]
			end--
		}
	}
	var ret int64
	for i, n := range nums {
		if i&1 == 0 {
			ret += int64(n)
			continue
		}
		ret -= int64(n)
	}
	return ret
}
