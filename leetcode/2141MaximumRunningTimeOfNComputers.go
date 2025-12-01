package leetcode

import "sort"

func maxRunTime(n int, batteries []int) int64 {
	// 1. Sort the batteries.
	sort.Ints(batteries)

	// 2. Calculate the sum of all "extra" batteries (the smallest ones).
	// Use int64 for sum to avoid potential overflow, as batteries can be large.
	var extra int64
	for _, b := range batteries[:len(batteries)-n] {
		extra += int64(b)
	}

	// 3. 'live' batteries are the n largest batteries for the n computers.
	// We work directly with the last n elements of the sorted 'batteries' slice.
	// Convert to int64 for consistent arithmetic.
	live := make([]int64, n)
	for i := 0; i < n; i++ {
		live[i] = int64(batteries[len(batteries)-n+i])
	}

	// 4. Iterate through the 'live' batteries, attempting to "level up" the runtimes.
	// This process effectively calculates how long the smallest battery (live[i])
	// can run until it catches up to the next smallest battery (live[i+1])
	// using the 'extra' power.

	for i := 0; i < n-1; i++ {
		// i+1 is the number of computers we are currently leveling up (live[0] to live[i]).

		// availableIncrease is the max increase we can apply to the current group
		// of (i+1) smallest batteries using the remaining 'extra' power.
		availableIncrease := extra / int64(i+1)

		// requiredIncrease is the power difference needed to level the smallest
		// group (up to live[i]) to the next largest battery (live[i+1]).
		requiredIncrease := live[i+1] - live[i]

		// If the 'extra' power, when distributed among the current group (i+1 computers),
		// is less than the required difference to reach live[i+1],
		// then the final max runtime must be between live[i] and live[i+1].
		if availableIncrease < requiredIncrease {
			return live[i] + availableIncrease
		}

		// If we have enough power, we use the required power to level them up.
		// Reduce 'extra' by the total power used for this step:
		// (number of computers) * (increase amount)
		extra -= int64(i+1) * requiredIncrease
	}

	// 5. If we reach the end of the loop, all 'live' batteries are leveled up
	// to the value of the largest one (live[n-1]). Any remaining 'extra' power
	// can be distributed equally among all 'n' computers.
	return live[n-1] + extra/int64(n)
}
