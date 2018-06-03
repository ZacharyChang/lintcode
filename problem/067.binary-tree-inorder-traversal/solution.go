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

var (
	result = make([]int, 0)
)

/**
 * Recursive
 * @param root: A Tree
 * @return: Postorder in ArrayList which contains node values.
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	inorderTraversal(root.Left)
	result = append(result, root.Val)
	inorderTraversal(root.Right)
	return result
}
