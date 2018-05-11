public class Solution {

    /*
     * @param root: The root of tree
     * 
     * @return: the head of doubly list node
     */
    public DoublyListNode bstToDoublyList(TreeNode root) {
        if (root == null) {
            return null;
        }
        DoublyListNode left = bstToDoublyList(root.left);
        DoublyListNode node = new DoublyListNode(root.val);
        DoublyListNode right = bstToDoublyList(root.right);

        node.next = right;
        if (right != null) {
            right.prev = node;
        }
        DoublyListNode cur = left;
        if (left != null) {
            while (cur.next != null) {
                cur = cur.next;
            }
            cur.next = node;
            node.prev = cur;
            node = left;
        }
        return node;
    }

    public static void main(String[] args) {
        TreeNode root = new TreeNode(4);
        root.left = new TreeNode(2);
        root.right = new TreeNode(5);
        DoublyListNode cur = new Solution().bstToDoublyList(root);
        DoublyListNode last = cur;
        while (cur != null) {
            last = cur;
            System.out.print(cur.val + "->");
            cur = cur.next;
        }
        System.out.println();
        while (last != null) {
            System.out.print(last.val + "->");
            last = last.prev;
        }
    }
}