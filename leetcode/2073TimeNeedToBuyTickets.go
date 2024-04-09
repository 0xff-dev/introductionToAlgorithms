package leetcode

func timeRequiredToBuy(tickets []int, k int) int {
	time := 0
	for i := 0; i < len(tickets); i++ {
		if i <= k {
			// 几轮
			time += min(tickets[i], tickets[k])
		} else {
			time += min(tickets[k]-1, tickets[i])
		}
	}
	return time
}
