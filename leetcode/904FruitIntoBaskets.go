package leetcode

type treeType struct {
	fruitType, count int
}

// 这道题
func totalFruit(fruits []int) int {
	trees := []treeType{{fruitType: fruits[0], count: 1}}
	idx, pre := 0, fruits[0]
	for i := 1; i < len(fruits); i++ {
		if fruits[i] == pre {
			trees[idx].count++
			continue
		}
		pre = fruits[i]
		trees = append(trees, treeType{fruitType: pre, count: 1})
		idx++
	}

	// 一种果树，选择全部
	if len(trees) == 1 {
		return trees[0].count
	}
	ans := -1
	a, b := trees[0], trees[1]
	sum := a.count + b.count
	start := 2
	// 1,2,1,2, 3
	for ; start < len(trees); start++ {
		if trees[start].fruitType == a.fruitType || trees[start].fruitType == b.fruitType {
			a, b = b, trees[start]
			sum += trees[start].count
			continue
		}
		if sum > ans {
			ans = sum
		}
		a, b = trees[start-1], trees[start]
		sum = a.count + b.count
	}
	if sum > ans {
		ans = sum
	}
	return ans
}
