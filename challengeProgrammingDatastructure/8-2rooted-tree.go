package challengeprogrammingdatastructure

import (
	"fmt"
)

func findRoot(tree []treeDefine) int {
	nodes := make([]bool, len(tree))
	for _, item := range tree {
		for _, ch := range item.child {
			nodes[ch] = true
		}
	}
	for i := 0; i < len(tree); i++ {
		if !nodes[i] {
			return i
		}
	}
	return -1
}

func RootedTree(tree []treeDefine) {
	father := make([]int, len(tree))
	depth := make([]int, len(tree))
	treeInfo := make([]treeDefine, len(tree))
	for _, item := range tree {
		treeInfo[item.id] = item
	}

	fatherNode := findRoot(tree)
	father[fatherNode] = -1
	var dfs func(int, int)
	dfs = func(start, dep int) {
		depth[start] = dep
		for _, ch := range treeInfo[start].child {
			father[ch] = start
			dfs(ch, dep+1)
		}
	}
	dfs(fatherNode, 0)
	for i := 0; i < len(tree); i++ {
		fmt.Printf("node %d: parent = %d, depth = %d, ", i, father[i], depth[i])
		if i == fatherNode {
			fmt.Print("root, ")
		} else if treeInfo[i].k == 0 {
			fmt.Print("leaf, []")
		} else {
			fmt.Print("internal node, ")
		}
		if treeInfo[i].k != 0 {
			fmt.Print("[")
			for j := 0; j < treeInfo[i].k; j++ {
				if j != 0 {
					fmt.Print(", ")
				}
				fmt.Print(treeInfo[i].child[j])
			}
			fmt.Print("]")
		}
		fmt.Println()
	}
}
