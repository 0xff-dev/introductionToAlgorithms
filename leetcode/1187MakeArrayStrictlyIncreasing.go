package leetcode

import "sort"

type pair1187 struct {
	x, y int
}

func makeArrayIncreasing(arr1 []int, arr2 []int) int {
	dp := make(map[pair1187]int)
	sort.Ints(arr2)

	var dfs func(int, int) int
	var bsearch func(int) int
	bsearch = func(val int) int {
		left, right := 0, len(arr1)
		for left < right {
			mid := (right-left)/2 + left
			if arr2[mid] <= val {
				left = mid + 1
			} else {
				right = mid
			}
		}
		return left
	}
	dfs = func(i, prev int) int {
		if i == len(arr1) {
			return 0
		}
		if v, ok := dp[pair1187{i, prev}]; ok {
			return v
		}

		cost := 2001
		if arr1[i] > prev {
			cost = dfs(i+1, arr1[i])
		}

		idx := bsearch(prev)
		if idx < len(arr2) {
			if r := dfs(i+1, arr2[idx]); r+1 < cost {
				cost = r + 1
			}
		}
		dp[pair1187{i, prev}] = cost
		return cost
	}
	ans := dfs(0, -1)
	if ans < 2001 {
		return ans
	}
	return -1
}
