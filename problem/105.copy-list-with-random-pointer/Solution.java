import java.util.HashMap;
import java.util.Map;

public class Solution {
    private Map<RandomListNode, Integer> map = new HashMap<>();

    /**
     * @param head: The head of linked list with a random pointer.
     * @return: A new head of a deep copy of the list.
     */
    public RandomListNode copyRandomList(RandomListNode head) {
        if (head == null) {
            return null;
        }
        if(map.isEmpty()){
            RandomListNode cur = head;
            while (cur != null) {
                map.put(cur, cur.label);
                cur = cur.next;
            }
        }
        RandomListNode node = new RandomListNode(head.label);
        node.next = copyRandomList(head.next);
        if (head.random != null) {
            node.random = new RandomListNode(map.get(head.random));
        } else {
            node.random = null;
        }
        return node;
    }
}