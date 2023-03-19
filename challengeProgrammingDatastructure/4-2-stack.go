package challengeprogrammingdatastructure

import "strconv"

func ReversePolish(expr []string) int {
	ans := 0
	if len(expr) == 0 {
		return ans
	}
	stack := make([]int, len(expr))
	index := 0
	for _, exp := range expr {
		if exp != "+" && exp != "-" && exp != "*" {
			v, _ := strconv.Atoi(exp)
			stack[index] = v
			index++
			continue
		}
		if index < 2 {
			return 0
		}
		a, b := stack[index-2], stack[index-1]
		index -= 2
		if exp == "+" {
			stack[index] = a + b
		}
		if exp == "-" {
			stack[index] = a - b
		}
		if exp == "*" {
			stack[index] = a * b
		}
		index++
	}
	return stack[0]
}
