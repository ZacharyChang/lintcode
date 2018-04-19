package lintcode

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * @param head: n
 * @return: The new head of reversed linked list.
 */
func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var array []*ListNode
	cur := head
	for cur != nil {
		array = append(array, cur)
		cur = cur.Next
	}
	for i := len(array) - 1; i > 0; i-- {
		array[i].Next = array[i-1]
	}
	array[0].Next = nil
	return array[len(array)-1]
}
