package leetcode

func smallestUniqueSubarray(nums []int) int {
	n := len(nums)

	sa := make([]int, n)
	rank := make([]int, n)
	nextRank := make([]int, n)

	maxRank := 0
	for i, num := range nums {
		sa[i] = i
		rank[i] = num
		maxRank = max(maxRank, num)
	}

	// Construct the sorted suffix array by doubling
	temp := make([]int, n)
	cnt := make([]int, max(n+1, maxRank+2))

	countingSort := func(getKey func(int) int) {
		clear(cnt[:maxRank+2])

		for _, i := range sa {
			cnt[getKey(i)]++
		}

		for i := 1; i <= maxRank+1; i++ {
			cnt[i] += cnt[i-1]
		}

		for p := n - 1; p >= 0; p-- {
			i := sa[p]
			key := getKey(i)
			cnt[key]--
			temp[cnt[key]] = i
		}

		sa, temp = temp, sa
	}

	for k := 1; k < n; k <<= 1 {
		// Sort by the second key, then stably sort by the first key
		countingSort(func(i int) int {
			if i+k < n {
				return rank[i+k] + 1
			}
			return 0
		})
		countingSort(func(i int) int {
			return rank[i] + 1
		})

		nextRank[sa[0]] = 0

		// Assign ranks to the newly sorted suffixes
		for i := 1; i < n; i++ {
			x, y := sa[i-1], sa[i]
			rx, ry := -1, -1
			if x+k < n {
				rx = rank[x+k]
			}
			if y+k < n {
				ry = rank[y+k]
			}

			nextRank[y] = nextRank[x]
			if rank[x] != rank[y] || rx != ry {
				nextRank[y]++
			}
		}

		rank, nextRank = nextRank, rank
		maxRank = rank[sa[n-1]]

		// All n suffixes have distinct ranks, so sorting is complete
		if maxRank == n-1 {
			break
		}
	}

	// Construct LCP using Kasai's algorithm
	pos := make([]int, n)
	lcp := make([]int, n)
	for i, index := range sa {
		pos[index] = i
	}

	h := 0
	for i := 0; i < n; i++ {
		if pos[i] == 0 {
			h = 0
			continue
		}

		j := sa[pos[i]-1]
		for i+h < n && j+h < n && nums[i+h] == nums[j+h] {
			h++
		}

		lcp[pos[i]] = h

		if h > 0 {
			h--
		}
	}

	answer := n
	for i := 0; i < n; i++ {
		next := 0
		if i+1 < n {
			next = lcp[i+1]
		}

		need := max(lcp[i], next) + 1
		if sa[i]+need <= n {
			answer = min(answer, need)
		}
	}

	return answer
}
