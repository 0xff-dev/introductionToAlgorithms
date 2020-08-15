package offer


func FindNumAppearOnce(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	_num := nums[0]
	for index := 1; index < len(nums); index++ {
		_num ^= nums[index]
	}
	return _num
}

// 有两个不同的数字，检测出来

func IsBitOne(num, index int) bool {
	num = num >> index
	return num & 1 == 1
}

func FindFirstBit1(num int) int {
	// return index
	index := 0
	for num & 1 == 0 {
		index++
		num >>= 1
	}
	return index
}

func TwoDiffNum(nums []int) (int, int) {
	index := FindFirstBit1(FindNumAppearOnce(nums))
	num1, num2 := 0, 0
	for _, _num := range nums {
		if IsBitOne(_num, index) {
			num1 ^= _num
		} else {
			num2 ^= _num
		}
	}
	return num1, num2
}

