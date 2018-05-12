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
 * @param inorder: A list of integers that inorder traversal of a tree
 * @param postorder: A list of integers that postorder traversal of a tree
 * @return: Root of a tree
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	return build(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1)
}

func build(inorder []int, postorder []int, inLeft int, inRight int, postLeft int, postRight int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	// 注意特殊情况
	if inRight < 0 || postRight < 0 || inLeft > inRight || postLeft > postRight {
		return nil
	}
	if inLeft == inRight {
		return &TreeNode{
			inorder[inLeft],
			nil,
			nil,
		}
	}
	if postLeft == postRight {
		return &TreeNode{
			postorder[postRight],
			nil,
			nil,
		}
	}
	rootVal := postorder[postRight]
	root := &TreeNode{
		rootVal,
		nil,
		nil,
	}
	rootLoc := 0
	for i := 0; i < len(inorder); i++ {
		if rootVal == inorder[i] {
			rootLoc = i
			break
		}
	}
	// 根据中序遍历得到左子树（或右子树）节点个数，由此推算后续遍历中的左子树和右子树下标
	leftNum := rootLoc - inLeft
	// rightNum := inRight - rootLoc
	// 中序下标： [inLeft:rootLoc-1],rootLoc(根节点),[rootLoc+1:inRight]
	// 后序下标： [postLeft:postLeft+左子树节点个数],[postLeft+左子树节点个数+1:postRight-1],postRight(根节点)
	root.Left = build(inorder, postorder, inLeft, rootLoc-1, postLeft, postLeft+leftNum-1)
	root.Right = build(inorder, postorder, rootLoc+1, inRight, postLeft+leftNum, postRight-1)
	return root
}
