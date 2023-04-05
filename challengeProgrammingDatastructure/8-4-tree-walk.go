package challengeprogrammingdatastructure

import "fmt"

func findFather84(tree []tree84) int {
	nodes := make([]bool, len(tree))
	for _, item := range tree {
		if item.left != -1 {
			nodes[item.left] = true
		}
		if item.right != -1 {
			nodes[item.right] = true
		}
	}
	for i := 0; i < len(nodes); i++ {
		if !nodes[i] {
			return i
		}
	}
	return -1
}
func treePreOrder(node int, treeMap []tree84) {
	fmt.Print(node, " ")
	if treeMap[node].left != -1 {
		treePreOrder(treeMap[node].left, treeMap)
	}
	if treeMap[node].right != -1 {
		treePreOrder(treeMap[node].right, treeMap)
	}
}
func treeInOrder(node int, treeMap []tree84) {
	if treeMap[node].left != -1 {
		treeInOrder(treeMap[node].left, treeMap)
	}
	fmt.Print(node, " ")
	if treeMap[node].right != -1 {
		treeInOrder(treeMap[node].right, treeMap)
	}
}

func treePostOrder(node int, treeMap []tree84) {
	if treeMap[node].left != -1 {
		treePostOrder(treeMap[node].left, treeMap)
	}
	if treeMap[node].right != -1 {
		treePostOrder(treeMap[node].right, treeMap)
	}
	fmt.Print(node, " ")
}

func TreeWalk(tree []tree84) {
	root := findFather84(tree)
	if root == -1 {
		return
	}
	treeMap := make([]tree84, len(tree))
	for _, item := range tree {
		treeMap[item.id] = item
	}

	fmt.Println("Preorder")
	treePreOrder(root, treeMap)
	fmt.Println()
	fmt.Println("Inorder")
	treeInOrder(root, treeMap)
	fmt.Println()
	fmt.Println("Postorder")
	treePostOrder(root, treeMap)
	fmt.Println()
}
