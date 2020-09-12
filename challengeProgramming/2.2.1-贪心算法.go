// 贪心算法就是遵循某种规则，选取当前认为是最优的解
package challengeProgramming

/*
	由1元，5元，10元，50元，100元，500元的硬币各n1，n2, n3枚，利用这些硬币支付A元，
	最少需要多少硬币。假定至少存在一种支付方案
*/

func MinCoins(money int, coins []int) int {
	if len(coins) != 6 {
		return 0
	}
	faceVal := map[int]int{
		5: 500, 4: 100, 3: 50,
		2: 10, 1: 5, 0: 1,
	}
	cnt, coinIndex := 0, 5
	for money > 0 && coinIndex > -1 {
		if coins[coinIndex] == 0 {
			coinIndex--
			continue
		}
		need := money / faceVal[coinIndex]
		remaining := money % faceVal[coinIndex]
		if need > coins[coinIndex] {
			need = coins[coinIndex]
			remaining = money - coins[coinIndex]*faceVal[coinIndex]
		}
		cnt, money = cnt+need, remaining
		coinIndex--
	}
	if money > 0 {
		return -1
	}
	return cnt
}
