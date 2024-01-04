package leetcode

func minOperations2870(nums []int) int {
	pos := make(map[int]int)
	arr := make([]int, 0)
	m := 1
	index := 0
	for _, n := range nums {
		v, ok := pos[n]
		if !ok {
			pos[n] = index
			arr = append(arr, 1)
			index++
			continue
		}
		arr[v]++
		if arr[v] > m {
			m = arr[v]
		}
	}

	ans := 0
	for _, n := range arr {
		if n == 1 {
			return -1
		}
		three := n / 3
		mod := n % 3
		if mod == 1 {
			three-- //挪出去一个3，组成3+1
			three += 2
		}
		if mod == 2 {
			three++
		}
		ans += three
	}
	return ans
}
