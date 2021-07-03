package leetcode

import (
	"fmt"
	"sort"
)

func findClosestElements(arr []int, k int, x int) []int {
	if len(arr) < k {
		return arr
	}
	sort.SliceStable(arr, func(i, j int) bool {
		a, b := arr[i]-x, arr[j]-x
		if a < 0 {
			a = -a
		}
		if b < 0 {
			b = -b
		}
		return a < b
	})
	fmt.Println(arr)
	sort.Ints(arr[:k])
	return arr[:k]
}

func s2(arr []int, k, x int) []int {
	l, r := 0, len(arr)
	if r <= k {
		return arr
	}
	if arr[0] >= x {
		return arr[:k]
	}
	if x >= arr[r-1] {
		return arr[r-k:]
	}

	for r-l > k {
		a, b := arr[l]-x, arr[r-1]-x
		if a < 0 {
			a = -a
		}
		if b < 0 {
			b = -b
		}
		if a <= b {
			r--
			continue
		}
		l++
	}
	return arr[l:r]
}
