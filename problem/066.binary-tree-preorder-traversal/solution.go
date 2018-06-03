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
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result = append(result, root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
	return result
}
