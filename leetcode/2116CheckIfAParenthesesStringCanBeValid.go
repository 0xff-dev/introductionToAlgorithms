package leetcode

/*
func canBeValid(s string, locked string) bool {
	stack := make([]int, len(s))
	index := -1
	for i, b := range s {
		if b == '(' {
			index++
			stack[index] = i
			continue
		}
		if index == -1 || stack[index] == ')' {
			index++
			stack[index] = i
			continue
		}
		index--
	}

	if index == -1 {
		return true
	}
	if index&1 == 0 {
		return false
	}

	var (
		checkL func(start, end int, count [][2]int) int
		checkR func(start, end int, count [][2]int) int
	)
	checkL = func(start, end int, count [][2]int) int {
		count[end] = [2]int{0, 0}
		minIndex := -1
		if locked[stack[end]] == '0' {
			count[end][0]++
		} else {
			count[end][1]++
			minIndex = end
		}
		for i := end - 1; i >= start; i-- {
			if locked[stack[i]] == '0' {
				count[i] = [2]int{count[i+1][0] + 1, count[i+1][1]}
			} else {
				count[i] = [2]int{count[i+1][0], count[i+1][1] + 1}
				minIndex = i
			}
		}
		return minIndex
	}

	checkR = func(start, end int, count [][2]int) int {
		count[start] = [2]int{0, 0}
		maxIndex := -1
		if locked[stack[start]] == '0' {
			count[end][0]++
		} else {
			count[start][1]++
			maxIndex = start
		}
		for i := start + 1; i <= end; i++ {
			if locked[stack[i]] == '0' {
				count[i] = [2]int{count[i-1][0] + 1, count[i-1][1]}
			} else {
				count[i] = [2]int{count[i-1][0], count[i-1][1] + 1}
				maxIndex = i
			}
		}
		return maxIndex
	}
	count := make([][2]int, index+1)
	if s[stack[0]] == '(' {
		minIndex := checkL(0, index, count)
		if minIndex != -1 {
			if count[minIndex][1] > count[minIndex][0] {
				return false
			}
		}

		return (count[0][0]-count[0][1])%1 == 0
	}
	if s[stack[index]] == ')' {
		maxIndex := checkR(0, index, count)
		if maxIndex != -1 {
			if count[maxIndex][1] > count[maxIndex][0] {
				return false
			}
		}
		return (count[index][0]-count[index][1])&1 == 0

	}

	split := -1
	for i := 0; i <= index; i++ {
		if s[stack[i]] == '(' {
			split = i
			break
		}
	}

	maxIndex := checkR(0, split-1, count)
	if maxIndex != -1 {
		if count[split-1][1] > count[split-1][0] {
			return false
		}
	}
	diff1 := count[split-1][0] - count[split-1][1]
	minIndex := checkL(split, index, count)
	if minIndex != -1 {
		if count[minIndex][1] > count[minIndex][0] {
			return false
		}
	}
	diff2 := count[split][0] - count[split][1]
	return (diff1+diff2)&1 == 0
}
*/

func canBeValid(s string, locked string) bool {
	l := len(s)
	if l&1 == 1 {
		return false
	}
	stack1, stack2 := make([]int, l), make([]int, l)
	i1, i2 := -1, -1
	for i, b := range s {
		if locked[i] == '0' {
			i1++
			stack1[i1] = i
			continue
		}
		if b == '(' {
			i2++
			stack2[i2] = i
			continue
		}
		if i2 != -1 {
			i2--
			continue
		}
		if i1 != -1 {
			i1--
			continue
		}
		return false
	}
	for i2 >= 0 && i1 >= 0 {
		if stack2[i2] > stack1[i1] {
			break
		}
		i2, i1 = i2-1, i1-1
	}
	return i2 == -1
}
