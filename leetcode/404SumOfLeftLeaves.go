package leetcode

func sumOfLeftLeaves(root *TreeNode) int {
    sum := 0
    tmp(root, &sum)
    return sum
}

func tmp(root *TreeNode, sum *int) {
    if root == nil {
        return
    }

    if root.Left != nil {
        left := root.Left
        if left.Left == nil && left.Right == nil {
            *sum+= left.Val
        } else {
            tmp(left, sum)
        }
    }
    if root.Right != nil {
        tmp(root.Right, sum)
    }
}
