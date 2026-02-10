package leetcode

func longestBalanced(nums []int) int {
	maxLen := 0
	n := len(nums)

	for left := 0; left < n; left++ {
		evenSeen := make(map[int]struct{})
		oddSeen := make(map[int]struct{})

		for right := left; right < n; right++ {
			num := nums[right]

			if num%2 == 0 {
				evenSeen[num] = struct{}{}
			} else {
				oddSeen[num] = struct{}{}
			}

			if len(evenSeen) == len(oddSeen) {
				currentLen := right - left + 1
				if currentLen > maxLen {
					maxLen = currentLen
				}
			}
		}
	}

	return maxLen
}
