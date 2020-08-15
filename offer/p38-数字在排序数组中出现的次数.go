package offer

import "log"

func binarySearch(nums []int, k int) int {
	start, end := 0, len(nums)-1
	for start < end {
		mid := (end-start)/2+start
		if nums[mid] == k {
			return mid
		} else if nums[mid] < k {
			start = mid+1
		} else {
			end = mid-1
		}
	}
	return -1
}

// O(n)+O(log n)
func existsKinArray(nums []int, k int) int {
	index := binarySearch(nums, k)
	if index == -1 {
		log.Println("not found ", k)
		return index
	}
	cnt := 0
	left, right := index, index+1
	for ; left >= 0 && nums[left] == k; left--{
		cnt++
	}
	for ; right < len(nums) && nums[right] == k; right++ {
		cnt++
	}
	return cnt
}
