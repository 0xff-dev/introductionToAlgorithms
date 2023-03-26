package challengeprogrammingdatastructure

import "fmt"

type tree83 struct {
	id          int
	left, right int
}

func findFather83(tree []tree83) int {
	nodes := make([]bool, len(tree))
	for _, item := range tree {
		if item.left != -1 {
			nodes[item.left] = true
		}
		if item.right != -1 {
			nodes[item.right] = true
		}
	}
	for i := 0; i < len(tree); i++ {
		if !nodes[i] {
			return i
		}
	}
	return -1
}

func dfs83(node, nodeDep int, treeMap []tree83, depth, height, father, sibling []int) int {
	n := treeMap[node]
	depth[node] = nodeDep
	if n.left == -1 && n.right == -1 {
		return 0
	}

	sibling[n.left] = -1
	sibling[n.right] = -1
	if n.left != -1 && n.right != -1 {
		sibling[n.left] = n.right
		sibling[n.right] = n.left
	}

	h := 0
	if n.left != -1 {
		father[n.left] = node
		h1 := dfs83(n.left, nodeDep+1, treeMap, depth, height, father, sibling)
		if h1 > h {
			h = h1
		}
	}
	if n.right != -1 {
		father[n.right] = node
		h1 := dfs83(n.right, nodeDep+1, treeMap, depth, height, father, sibling)
		if h1 > h {
			h = h1
		}
	}
	h++
	height[node] = h
	return h
}
func BinaryTreePritn(tree []tree83) {
	root := findFather83(tree)
	if root == -1 {
		return
	}
	treeMap := make([]tree83, len(tree))
	for _, item := range tree {
		treeMap[item.id] = item
	}
	depth := make([]int, len(tree))
	height := make([]int, len(tree))
	father := make([]int, len(tree))
	father[root] = -1
	silbing := make([]int, len(tree))
	silbing[root] = -1
	dfs83(root, 0, treeMap, depth, height, father, silbing)
	for i := 0; i < len(tree); i++ {
		fmt.Printf("node %d: parent = %d, sibling = %d, degree = ", i, father[i], silbing[i])
		degree := 0
		if treeMap[i].left != -1 {
			degree++
		}
		if treeMap[i].right != -1 {
			degree++
		}
		fmt.Print(degree, ", ")
		fmt.Printf("depth = %d, height = %d, ", depth[i], height[i])
		nodeType := "internal node"
		if treeMap[i].left == -1 && treeMap[i].right == -1 {
			nodeType = "leaf"
		}
		if i == root {
			nodeType = "root"
		}
		fmt.Print(nodeType, "\n")
	}
}
