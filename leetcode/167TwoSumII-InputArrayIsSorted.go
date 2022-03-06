package leetcode

func twoSum167(numbers []int, target int) []int {
	start, end := 0, len(numbers)-1
	for start < end {
		//mid := (end-start)/2 + start
		sum := numbers[start] + numbers[end]
		if sum == target {
			return []int{start + 1, end + 1}
		}
		if sum < target {
			start++
			continue
		}
		end--
	}
	return nil
}

func twoSum_1(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		low, high := i+1, len(numbers)-1
		for low <= high {
			mid := (high-low)/2 + low
			if numbers[mid] == target-numbers[i] {
				return []int{i + 1, mid + 1}
			} else if numbers[mid] > target-numbers[i] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return []int{-1, -1}
}
