package _3_merge_k_sorted_lists

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//两两合并
func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if l == 0 {
		return nil
	}
	if l == 1 {
		return lists[0]
	}
	for i:=0;i<l/2;i++{
		lists[i] = mergeTwoLists(lists[i], lists[l-1-i])
	}
	return mergeKLists(lists[:(l+1)/2])
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
