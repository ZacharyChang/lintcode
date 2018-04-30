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

var treeMap map[int]int

/**
 * @param inOrder: A list of integers that inorder traversal of a tree
 * @param preOrder: A list of integers that preorder traversal of a tree
 * @return: Root of a tree
 */
func buildTree(inOrder []int, preOrder []int) *TreeNode {
	treeMap = make(map[int]int, len(inOrder))
	for i := 0; i < len(inOrder); i++ {
		treeMap[inOrder[i]] = i
	}
	return build(inOrder, preOrder, 0, len(inOrder)-1, 0, len(preOrder)-1)
}

func build(inOrder []int, preOrder []int, inLeft int, inRight int, preLeft int, preRight int) *TreeNode {
	if len(inOrder) == 0 || len(preOrder) == 0 {
		return nil
	}
	// 注意特殊情况
	if inRight < 0 || preRight < 0 || inLeft > inRight || preLeft > preRight {
		return nil
	}
	rootVal := preOrder[preLeft]
	root := &TreeNode{
		rootVal,
		nil,
		nil,
	}
	rootLoc := treeMap[rootVal]
	if inLeft == inRight || preLeft == preRight {
		return root
	}
	// 根据中序遍历得到左子树（或右子树）节点个数，由此推算先续遍历中的左子树和右子树下标
	leftNum := rootLoc - inLeft
	// rightNum := inRight - rootLoc
	// 中序下标： [inLeft:rootLoc-1],rootLoc(根节点),[rootLoc+1:inRight]
	// 先序下标： preLeft(根节点）,[preLeft+1:preLeft+1+(左子树节点个数-1)],[preLeft+1+(左子树节点个数-1)+1:preRight]
	root.Left = build(inOrder, preOrder, inLeft, rootLoc-1, preLeft+1, preLeft+leftNum)
	root.Right = build(inOrder, preOrder, rootLoc+1, inRight, preLeft+leftNum+1, preRight)
	return root
}
