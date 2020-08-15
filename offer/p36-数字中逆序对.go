package offer

import "fmt"

func InverseParis(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func merge(start, end int, nums []int) int {
	inversePair := 0
	mid := (end-start)/2 + start
	i, j, index := start, mid+1, 0
	result := make([]int, end-start+1)
	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			result[index] = nums[i]
			i++
		} else {
			result[index] = nums[j]
			for k := i; k <= mid; k++ {
				fmt.Printf("parir{%d:%d}\n", nums[k], nums[j])
			}
			inversePair += mid - i + 1
			j++
		}
		index++
	}
	for ; i <= mid; i++ {
		result[index] = nums[i]
		index++
	}
	for ; j <= end; j++ {
		result[index] = nums[j]
		index++
	}
	for k := 0; k < len(result); k++ {
		nums[start+k] = result[k]
	}
	return inversePair
}

func mergeSort(nums []int, start, end int) int {
	if start < end {
		mid := (end-start)/2 + start
		return mergeSort(nums, start, mid) + mergeSort(nums, mid+1, end) + merge(start, end, nums)
	}
	return 0
}
