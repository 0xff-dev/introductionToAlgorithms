package offer

func FindNumberWithSum(nums []int, sum int) (int, int, bool) {
	start, end := 0, len(nums)-1
	for start < end {
		tmpSum := nums[start] + nums[end]
		if tmpSum == sum {
			return start, end, true
		}
		if tmpSum < sum {
			start++
		} else {
			end--
		}
	}
	return -1, -1, false
}

// 1 2 3 4 5 6 7 8
func FindMultiNumberWithSum(nums []int, sum int) [][]int {
	result := make([][]int, 0)
	tmpSum := 0
	start, end := 0, 0
	for index, num := range nums {
		tmpSum += num
		end = index
		if tmpSum == sum {
			result = append(result, nums[start:end+1])
			continue
		}
		for tmpSum > sum && start < end {
			tmpSum -= nums[start]
			start++
			if tmpSum == sum {
				result = append(result, nums[start:end+1])
				break
			}
		}
	}
	return result
}
