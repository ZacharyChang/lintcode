import java.util.List;
import java.util.ArrayList;
import java.util.Arrays;

public class Solution {
    private List<List<Integer>> result;

    /*
     * @param root: the root of binary tree
     * 
     * @param target: An integer
     * 
     * @return: all valid paths
     */
    public List<List<Integer>> binaryTreePathSum(TreeNode root, int target) {
        result = new ArrayList<List<Integer>>();
        List<Integer> currentPath = new ArrayList<Integer>();
        binaryTreePathSum(root, currentPath, 0, target);
        return result;
    }

    public List<List<Integer>> binaryTreePathSum(TreeNode cur, List<Integer> currentPath, int sum, int target) {
        if (cur == null) {
            return result;
        }
        currentPath.add(cur.val);
        sum += cur.val;
        // only the path from root to leaf nodes is available
        if (sum == target && cur.left == null && cur.right == null) {
            result.add(new ArrayList<>(currentPath));
        }
        binaryTreePathSum(cur.left, new ArrayList<>(currentPath), sum, target);
        binaryTreePathSum(cur.right, new ArrayList<>(currentPath), sum, target);
        return result;
    }

    public static void main(String[] args) {
        TreeNode root = new TreeNode(1);
        TreeNode firstChildLeft = new TreeNode(2);
        firstChildLeft.left = new TreeNode(2);
        firstChildLeft.right = new TreeNode(3);
        root.left = firstChildLeft;
        root.right = new TreeNode(4);
        List<List<Integer>> result = new Solution().binaryTreePathSum(root, 5);
        System.out.println(result);
    }
}