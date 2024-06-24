package leetcode

func minKBitFlips(nums []int, k int) int {
	// Keeps track of flipped states
	flipped := make([]bool, len(nums))
	// Tracks valid flips within the past window
	validFlipsFromPastWindow := 0
	// Counts total flips needed
	flipCount := 0

	for i := 0; i < len(nums); i++ {
		if i >= k {
			// Decrease count of valid flips from the past window if needed
			if flipped[i-k] {
				validFlipsFromPastWindow--
			}
		}

		// Check if current bit needs to be flipped
		if validFlipsFromPastWindow%2 == nums[i] {
			// If flipping the window extends beyond the array length,
			// return -1
			if i+k > len(nums) {
				return -1
			}
			// Increment the count of valid flips and mark current as flipped
			validFlipsFromPastWindow++
			flipped[i] = true
			flipCount++
		}
	}

	return flipCount
}

/*
func minKBitFlips(nums []int, k int) int {
		idx := 0
		ans := 0
		for idx < len(nums) {
			if nums[idx] == 0 {
				if idx+k > len(nums) {
					return -1
				}
				ans++
				for j := idx; j < idx+k; j++ {
					nums[j] = 1 - nums[j]
				}
				continue
			}
			idx++
		}
		return ans

}
*/
