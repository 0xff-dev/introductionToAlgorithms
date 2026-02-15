package leetcode

import (
	"strconv"
)

type NestedInteger385 struct {
}

func (n NestedInteger385) IsInteger() bool {
	return true
}
func (n NestedInteger385) GetInteger() int {
	return 0
}

func (n *NestedInteger385) SetInteger(value int) {}

func (n *NestedInteger385) Add(elem NestedInteger385) {}

func (n NestedInteger385) GetList() []*NestedInteger {
	return nil
}

func deserialize(s string) *NestedInteger {
	if len(s) == 0 {
		return nil
	}

	root := &NestedInteger{}
	if s[0] != '[' {
		num, _ := strconv.Atoi(s)
		root.SetInteger(num)
		return root
	}

	numStart := -1
	stack := []*NestedInteger{root}

	// "[123,[456,[789]]]"
	for i := 1; i < len(s); i++ {
		if s[i] == '-' || (s[i] >= '0' && s[i] <= '9') {
			if numStart == -1 {
				numStart = i
			}
			continue
		}
		if s[i] == ',' {
			if numStart != -1 {
				num, _ := strconv.Atoi(s[numStart:i])
				node := NestedInteger{}
				node.SetInteger(num)
				stack[len(stack)-1].Add(node)
				numStart = -1
			}
			continue
		}
		// root/ node1 /node2
		if s[i] == '[' {
			node := &NestedInteger{}
			stack = append(stack, node)
			continue
		}
		if s[i] == ']' {
			if numStart != -1 {
				num, _ := strconv.Atoi(s[numStart:i])
				node := NestedInteger{}
				node.SetInteger(num)
				stack[len(stack)-1].Add(node)
				numStart = -1
			}
			l := len(stack)
			top := stack[l-1]
			stack = stack[:l-1]
			if len(stack) > 0 {
				stack[len(stack)-1].Add(*top)
			}
		}
	}

	return root
}
