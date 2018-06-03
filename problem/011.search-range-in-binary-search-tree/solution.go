package lintcode

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * @param root: param root: The root of the binary search tree
 * @param k1: An integer
 * @param k2: An integer
 * @return: return: Return all keys that k1<=key<=k2 in ascending order
 */
func searchRange(root *TreeNode, k1 int, k2 int) []int {
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		if root.Val >= k1 && root.Val <= k2 {
			return []int{root.Val}
		}
		return []int{}
	}
	result := make([]int, 0)
	if root.Left != nil && root.Val > k1 {
		result = append(result, searchRange(root.Left, k1, k2)...)
	}
	if root.Val >= k1 && root.Val <= k2 {
		result = append(result, root.Val)
	}
	if root.Right != nil && root.Val < k2 {
		result = append(result, searchRange(root.Right, k1, k2)...)
	}
	return result
}
