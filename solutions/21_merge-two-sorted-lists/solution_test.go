package _1_merge_two_sorted_lists

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//recursion
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}

//normal
func mergeTwoListsV2(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	temp := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			temp.Next = l1
			l1 = l1.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
		}
		temp = temp.Next
	}
	if l1 == nil {
		temp.Next = l2
	} else {
		temp.Next = l1
	}

	return head.Next
}
