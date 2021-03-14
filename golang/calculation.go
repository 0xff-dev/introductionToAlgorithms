package golang

// +, -, *, / ()
func calculation(expression string) int {
	numStack, optStack := make([]int, 0), make([]byte, 0)
	idx := 0

	for idx < len(expression) {
		if expression[idx] == ' ' {
			idx++
			continue
		}
		if isOperator(expression[idx]) {
			if expression[idx] == ')' {
				topIdx := len(optStack) - 1
				for optStack[topIdx] != '(' {
					popStackCal(&numStack, &optStack)
					topIdx--
				}
				optStack = optStack[:topIdx] // delete (
				idx++
				continue
			}
			if len(optStack) == 0 || less(expression[idx], optStack[len(optStack)-1]) {
				optStack = append(optStack, expression[idx])
				idx++
				continue
			}
			length := len(optStack)
			for length > 0 && !less(expression[idx], optStack[length-1]) {
				if expression[idx] == '(' {
					break
				}
				popStackCal(&numStack, &optStack)
				length = len(optStack)
				if len(numStack) == 1 {
					break
				}
			}
			optStack = append(optStack, expression[idx])
			idx++
			continue
		}
		base := 0
		i := idx
		for ; i < len(expression) && isNum(expression[i]); i++ {
			base = base*10 + int(expression[i]-'0')
		}
		numStack = append(numStack, base)
		idx = i
	}
	for len(optStack) > 0 {
		popStackCal(&numStack, &optStack)
		if len(numStack) == 1 {
			break
		}
	}
	for len(numStack) > 0 {
		length := len(numStack)
		if length == 1 {
			return numStack[0]
		}
		topOne, topTwo := numStack[length-1], numStack[length-2]
		numStack = numStack[:length-2]
		numStack = append(numStack, topOne*topTwo)
	}
	return numStack[0]
}

func isOperator(b byte) bool {
	return b == '+' || b == '-' || b == '*' || b == '/' || b == '(' || b == ')'
}

func isNum(b byte) bool {
	return b >= '0' && b <= '9'
}

func less(in, top byte) bool {
	switch top {
	case '*', '/':
		return false
	case '+', '-':
		if in == '*' || in == '/' {
			return true
		}
		return false
	case '(':
		return true
	}
	return false
}

func numOpt(a, b int, opt byte) int {
	if opt == '+' {
		return a + b
	}
	if opt == '*' {
		return a * b
	}
	if opt == '/' {
		if b == 0 {
			panic("divide by zero")
		}
		return a / b
	}
	if opt == '-' {
		return a - b
	}
	return a * b // ab = a*b
}

func popStackCal(numStack *[]int, optStack *[]byte) {
	length, optStackLength := len(*numStack), len(*optStack)
	if length == 1 {
		return
	}
	topOne, topTwo := (*numStack)[length-1], (*numStack)[length-2]
	topOpt := (*optStack)[optStackLength-1]
	*numStack = (*numStack)[:length-2]
	*optStack = (*optStack)[:optStackLength-1]

	r := numOpt(topTwo, topOne, topOpt)
	*numStack = append(*numStack, r)
}
