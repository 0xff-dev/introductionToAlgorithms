package leetcode

func isSymmetric(root *TreeNode) bool {
    return isSymmetricAux(root, root)
}

func isSymmetricAux(t1, t2 *TreeNode) bool {
    if t1 == nil && t2 == nil {
        return true
    }

    if t1 == nil || t2 == nil {
        return false
    }

    return t1.Val == t2.Val && isSymmetricAux(t1.Left, t2.Right) && isSymmetricAux(t1.Right, t2.Left)
}
