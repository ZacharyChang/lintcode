public class Solution {
    /*
     * Two Pointers
     * @param headA: the first list
     * @param headB: the second list
     * @return: a ListNode
     */
    public ListNode getIntersectionNode(ListNode headA, ListNode headB) {
        ListNode pointA = headA;
        ListNode pointB = headB;
        ListNode intersectionNode = null;
        boolean loopA = true;
        boolean loopB = true;
        if(pointA==null || pointB==null){
            return null;
        }
        while(loopA || loopB){
            if(pointA.next==null){
                pointA = headB;
                loopA=false;
            }else{
                pointA=pointA.next;
            }
            if(pointB.next==null){
                pointB=headA;
                loopB=false;
            }else{
                pointB=pointB.next;
            }
        }
        while(pointA!=null || pointB!=null){
            if(pointA.val==pointB.val){
                if(intersectionNode==null){
                    intersectionNode=pointA;
                }
            }else{
                intersectionNode=null;
            }
            pointA=pointA.next;
            pointB=pointB.next;
        }
        return intersectionNode;
    }
}