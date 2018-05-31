public class SolutionTest{
    public static void printList(ListNode head){
        while(head!=null){
            System.out.print(head.val+"->");
            head=head.next;
        }
    }

    public static void main(String[] args){
        ListNode node1 = new ListNode(7);
        node1.next=new ListNode(9);
        node1.next.next=new ListNode(11);
        ListNode node2 = new ListNode(9);
        node2.next = new ListNode(11);
        printList(new Solution().getIntersectionNode(node1,node2));
    }
}