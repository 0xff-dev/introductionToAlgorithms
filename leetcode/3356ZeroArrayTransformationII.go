package leetcode

func minZeroArray(nums []int, queries [][]int) int {
	left, right := 0, len(queries)

	// Zero array isn't formed after all queries are processed
	if !canFormZeroArray(nums, queries, right) {
		return -1
	}

	// Binary Search
	for left <= right {
		middle := left + (right-left)/2
		if canFormZeroArray(nums, queries, middle) {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	// Return earliest query that zero array can be formed
	return left
}

func canFormZeroArray(nums []int, queries [][]int, k int) bool {
	n := len(nums)
	sum := 0
	differenceArray := make([]int, n+1)

	// Process query
	for queryIndex := 0; queryIndex < k; queryIndex++ {
		start := queries[queryIndex][0]
		end := queries[queryIndex][1]
		val := queries[queryIndex][2]

		// Process start and end of range
		differenceArray[start] += val
		if end+1 < len(differenceArray) {
			differenceArray[end+1] -= val
		}
	}

	// Check if zero array can be formed
	for numIndex := 0; numIndex < n; numIndex++ {
		sum += differenceArray[numIndex]
		if sum < nums[numIndex] {
			return false
		}
	}
	return true
}
