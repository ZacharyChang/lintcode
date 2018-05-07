package lintcode

/**
 * Definition for singly-linked list.
 *
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * @param head: The first node of linked list.
 * @param n: An integer
 * @return: The head of linked list.
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	begin := head
	cur := head
	// 多移动一位，便于移动到倒数第个节点的前一位
	for i := 0; i < n+1; i++ {
		// check the next
		if cur == nil {
			return head.Next
		}
		cur = cur.Next
	}
	for cur != nil {
		cur = cur.Next
		begin = begin.Next
	}
	begin.Next = begin.Next.Next
	return head
}
