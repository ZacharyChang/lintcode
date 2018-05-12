import java.util.HashMap;
import java.util.Map;

/**
 * Definition of TreeNode:
 * public class TreeNode {
 *     public int val;
 *     public TreeNode left, right;
 *     public TreeNode(int val) {
 *         this.val = val;
 *         this.left = this.right = null;
 *     }
 * }
 */
public class Solution {
    private Map<Integer,Integer> treeMap;

    /**
     * @param inorder: A list of integers that inorder traversal of a tree
     * @param postorder: A list of integers that postorder traversal of a tree
     * @return: Root of a tree
     */
    public TreeNode buildTree(int[] preorder, int[] inorder) {
        treeMap = new HashMap<>();
        for(int i=0;i<inorder.length;i++){
            treeMap.put(inorder[i],i);
        }
        return build(preorder,inorder,0,preorder.length-1,0,inorder.length-1);
    }

    public TreeNode build(int[] preorder, int[] inorder, int preLeft, int preRight, int inLeft, int inRight){
        if (preorder.length == 0||preLeft>preRight||inLeft>inRight){
            return null;
        }
        int rootVal = preorder[preLeft];
        TreeNode root = new TreeNode(rootVal);
        if (preLeft==preRight||inLeft==inRight){
            return root;
        }
        int rootLoc = treeMap.get(rootVal);
        int leftNum = rootLoc-inLeft;
        root.left=build(preorder, inorder, preLeft+1, preLeft+1+leftNum-1, inLeft, rootLoc-1);
        root.right=build(preorder, inorder, preLeft+1+leftNum, preRight, rootLoc+1, inRight);
        return root;
    }
}
