package dp

/*
	给一个字符串rrrryrryryrry 只包含ry两个字符，转换成rr....yyy.....rrr.的形式，请求最小的修改次数
*/
func ToRYR(leaves []byte) int {
	if len(leaves) < 3 {
		return -1
	}
	result := len(leaves)
	lr, lry, rr, rry := make([]int, result), make([]int, result), make([]int, result), make([]int, result)
	for index := 0; index < result; index++ {
		lr[index], lry[index], rr[index], rry[index] = result, result, result, result
	}
	countLY, countRY := 0, 0
	for index := 0; index < result; index++ {
		if leaves[index] == 'y' {
			countLY++
		}
		lr[index] = countLY
		if index > 0 {
			lry[index] = min(lry[index-1], lr[index-1])
			if leaves[index] == 'r' {
				lry[index]++
			}
		}
	}
	for index := result - 1; index >= 0; index-- {
		if leaves[index] == 'y' {
			countRY++
		}
		rr[index] = countRY
		if index < result-1 {
			rry[index] = min(rr[index+1], rry[index+1])
			if leaves[index] == 'r' {
				rry[index]++
			}
		}
	}

	for index := 1; index < len(leaves)-1; index++ {
		result = min(result, lr[index]+rry[index+1], lry[index]+rry[index+1], lry[index]+rr[index+1])
	}
	return result
}

func min(args ...int) int {
	m := args[0]
	for index := 1; index < len(args); index++ {
		if args[index] < m {
			m = args[index]
		}
	}
	return m
}
