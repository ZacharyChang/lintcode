package lintcode

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 数组暂存ListNode，然后改变Next指向
 * 需要额外消耗O(n)的空间
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

/**
 * 递归方法，不断递归直到head指向最后节点
 * 每次调转head与下一届点顺序，直到递归结束
 * 新链表的头在递归过程中不改变，依次向上返回
 * @param head: n
 * @return: The new head of reversed linked list.
 */
func solution2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	newHead := solution2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
