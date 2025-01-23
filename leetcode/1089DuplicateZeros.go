package leetcode

func duplicateZeros(arr []int) {
	tmp := make([]int, len(arr))
	index, i := 0, 0
	for ; i < len(arr) && index < len(arr); i++ {
		if arr[i] != 0 {
			tmp[index] = arr[i]
			index++
			continue
		}
		index += 2
	}
	copy(arr, tmp)
}
