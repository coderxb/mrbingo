package leetcode

// SliceToList 将切片转换为链表
func sliceToList(nums []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range nums {
		cur.next = &ListNode{val: v}
		cur = cur.next
	}
	return dummy.next
}

// ListToSlice 将链表转换为切片
func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.val)
		head = head.next
	}
	return result
}
