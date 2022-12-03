package leetcode

func lemonadeChange(bills []int) bool {
	change := [5]int{}
	for _, cost := range bills {
		idx := cost / 5

		if cost == 5 {
			change[idx]++
			continue
		}

		diff := cost - 5
		for change[2] > 0 && diff >= 10 {
			change[2]--
			diff -= 10
		}
		for change[1] > 0 && diff >= 5 {
			change[1]--
			diff -= 5
		}
		if diff > 0 {
			return false
		}
		change[idx]++
	}
	return true
}
