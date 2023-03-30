package challengeprogrammingdatastructure

func BinarySearchTreeSearch(root *BSTNode, target int) bool {
	for root != nil {
		if root.Val == target {
			return true
		}
		if root.Val < target {
			root = root.Right
			continue
		}
		root = root.Left
	}
	return false
}
