/**
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
**/

package leetcode

/* Definition for singly-linked list. */
type ListNode struct {
	val  int
	next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	current := dummyHead
	carry := 0

	if l1 == nil && l2 == nil {
		return dummyHead.next
	}
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry	
		if l1 != nil {
			sum += l1.val
			l1 = l1.next
		}
		if l2 != nil {
			sum += l2.val
			l2 = l2.next
		}
		carry = sum / 10
		current.next = &ListNode{val: sum % 10}
		current = current.next
	}
	return dummyHead.next
}
