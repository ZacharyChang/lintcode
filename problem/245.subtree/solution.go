package lintcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * @param T1: The roots of binary tree T1.
 * @param T2: The roots of binary tree T2.
 * @return: True if T2 is a subtree of T1, or false.
 */
func isSubtree(T1 *TreeNode, T2 *TreeNode) bool {
	// nil treenode is always a subtree
	if T2 == nil {
		return true
	}
	if T1 == nil {
		return false
	}
	if isSame(T1, T2) {
		return true
	}
	return isSubtree(T1.Left, T2) || isSubtree(T1.Right, T2)
}

func isSame(T1 *TreeNode, T2 *TreeNode) bool {
	if T1 == nil && T2 == nil {
		return true
	}
	if T1 == nil || T2 == nil {
		return false
	}
	if T1.Val == T2.Val && isSame(T1.Left, T2.Left) && isSame(T1.Right, T2.Right) {
		return true
	}
	return false
}
