package leetcode

import "sort"

func isPossibleDivide(nums []int, k int) bool {
	l := len(nums)
	if l%k != 0 {
		return false
	}
	if k == 1 {
		return true
	}

	// 1:1, 2:2, 3:2, 4:1, 6:1,7:1,8:1
	keys := make([]int, 0)
	keyCount := make(map[int]int)
	for _, n := range nums {
		keyCount[n]++
		if keyCount[n] == 1 {
			keys = append(keys, n)
		}
	}
	sort.Ints(keys)

	idx := 1
	count := keyCount[keys[0]]
	nextIdx := -1
	keyCount[keys[0]] = 0
	used := 1
	l -= count

	// 1:1, 2:2, 3:2, 4:1, 6:1,7:1,8:1
	for idx < len(keys) {
		if keys[idx]-keys[idx-1] != 1 {
			return false
		}
		keyCount[keys[idx]] -= count
		l -= count
		used++
		if keyCount[keys[idx]] < 0 {
			return false
		}
		if keyCount[keys[idx]] > 0 && nextIdx == -1 {
			nextIdx = idx
		}

		if used != k {
			idx++
			continue
		}
		if nextIdx == -1 {
			if idx == len(keys)-1 {
				break
			}
			idx++
			count = keyCount[keys[idx]]
			nextIdx = -1
			keyCount[keys[idx]] = 0
			l -= count
			used = 1
			idx++
			continue
		}

		idx = nextIdx + 1
		count = keyCount[keys[nextIdx]]
		keyCount[keys[nextIdx]] = 0
		nextIdx = -1
		l -= count
		used = 1
	}
	return l == 0 && used == k
}
