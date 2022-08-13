package leetcode

func numWaterBottles(numBottles int, numExchange int) int {
	emptyBottle := numBottles
	ans := numBottles
	for emptyBottle >= numExchange {
		bottles := emptyBottle / numExchange
		ans += bottles
		emptyBottle = emptyBottle%numExchange + bottles
	}
	return ans
}
