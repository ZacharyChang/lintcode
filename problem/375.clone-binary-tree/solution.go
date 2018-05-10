package lintcode

/**
 * Definition for a binary tree node.
 *
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * @param root: The root of binary tree
 * @return: root of new tree
 */
func cloneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return &TreeNode{
		root.Val,
		cloneTree(root.Left),
		cloneTree(root.Right),
	}
}
