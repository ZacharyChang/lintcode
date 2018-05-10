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
	result [][]int
	queue  []*TreeNode
)

/**
 * @param root: A Tree
 * @return: Level order a list of lists of integer
 */
func levelOrder(root *TreeNode) [][]int {
	result = make([][]int, 0)
	queue = make([]*TreeNode, 0)
	if root == nil {
		return result
	}
	queue = make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		queueCopy := make([]*TreeNode, len(queue))
		copy(queueCopy, queue)
		// clear the queue
		queue = make([]*TreeNode, 0)
		line := make([]int, 0)
		for _, node := range queueCopy {
			line = append(line, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, line)
	}
	return result
}
