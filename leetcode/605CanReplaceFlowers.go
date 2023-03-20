package leetcode

func canPlaceFlowers(flowerbed []int, n int) bool {
	can := 0
	zeroIndex := flowerbed[0]
	leftHasOne := flowerbed[0] == 1
	sum := 0

	for ; zeroIndex < len(flowerbed); zeroIndex++ {
		if flowerbed[zeroIndex] == 0 {
			sum++
			continue
		}
		sum -= 2
		if !leftHasOne {
			sum++
		}
		leftHasOne = true
		if sum <= 0 {
			sum = 0
			continue
		}

		can += sum / 2
		if sum&1 == 1 {
			can++
		}
		sum = 0
	}
	if sum != 0 {
		can += sum / 2
		if !leftHasOne && sum&1 == 1 {
			can++
		}
	}
	return can >= n
}
