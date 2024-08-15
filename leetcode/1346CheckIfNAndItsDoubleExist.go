package leetcode

import "sort"

func checkIfExist(arr []int) bool {
	sort.Ints(arr)
	l := len(arr)
	in := sort.Search(l, func(i int) bool {
		return arr[i] >= 0
	})

	for i := l - 1; i >= in; i-- {
		idx := sort.Search(i, func(ii int) bool {
			return arr[ii]*2 >= arr[i]
		})
		if idx != i && arr[idx]*2 == arr[i] {
			return true
		}
	}
	for i := in - 1; i > 0; i-- {
		idx := sort.Search(i, func(ii int) bool {
			return arr[i]*2 <= arr[ii]
		})
		if idx != i && arr[idx] == arr[i]*2 {
			return true
		}
	}

	return false
}
