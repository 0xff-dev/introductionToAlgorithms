package leetcode

var maxHeightAfterRemoval []int
var currentMaxHeight int

func treeQueries(root *TreeNode, queries []int) []int {
	maxHeightAfterRemoval = make([]int, 100001) // Adjust the size as needed
	currentMaxHeight = 0

	traverseLeftToRight(root, 0)
	currentMaxHeight = 0 // Reset for the second traversal
	traverseRightToLeft(root, 0)

	// Process queries and build the result slice
	queryResults := make([]int, len(queries))
	for i, query := range queries {
		queryResults[i] = maxHeightAfterRemoval[query]
	}

	return queryResults
}

// Left to right traversal
func traverseLeftToRight(node *TreeNode, currentHeight int) {
	if node == nil {
		return
	}

	// Store the maximum height if this node were removed
	maxHeightAfterRemoval[node.Val] = currentMaxHeight

	// Update the current maximum height
	currentMaxHeight = max(currentMaxHeight, currentHeight)

	// Traverse left subtree first, then right
	traverseLeftToRight(node.Left, currentHeight+1)
	traverseLeftToRight(node.Right, currentHeight+1)
}

// Right to left traversal
func traverseRightToLeft(node *TreeNode, currentHeight int) {
	if node == nil {
		return
	}

	// Update the maximum height if this node were removed
	maxHeightAfterRemoval[node.Val] = max(maxHeightAfterRemoval[node.Val], currentMaxHeight)

	// Update the current maximum height
	currentMaxHeight = max(currentHeight, currentMaxHeight)

	// Traverse right subtree first, then left
	traverseRightToLeft(node.Right, currentHeight+1)
	traverseRightToLeft(node.Left, currentHeight+1)
}
