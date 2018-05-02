import java.util.HashMap;
import java.util.Map;

public class Solution {
    private Map<Integer,Integer> treeMap;

    /**
     * @param inorder: A list of integers that inorder traversal of a tree
     * @param postorder: A list of integers that postorder traversal of a tree
     * @return: Root of a tree
     */
    public TreeNode buildTree(int[] inorder, int[] postorder) {
        treeMap = new HashMap<>();
        for(int i=0;i<inorder.length;i++){
            treeMap.put(inorder[i],i);
        }
        return build(inorder,postorder,0,inorder.length-1,0,postorder.length-1);
    }

    public TreeNode build(int[] inorder, int[] postorder, int inLeft, int inRight, int postLeft, int postRight){
        if (inorder.length == 0||postLeft>postRight||inLeft>inRight){
            return null;
        }
        int rootVal = postorder[postRight];
        TreeNode root = new TreeNode(rootVal);
        if (postLeft==postRight||inLeft==inRight){
            return root;
        }
        int rootLoc = treeMap.get(rootVal);
        int leftNum = rootLoc-inLeft;
        root.left=build(inorder, postorder, inLeft, rootLoc-1,postLeft,postLeft+leftNum-1);
        root.right=build(inorder, postorder, rootLoc+1, inRight, postLeft+leftNum,postRight-1 );
        return root;
    }
}
