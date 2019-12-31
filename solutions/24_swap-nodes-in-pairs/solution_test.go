package _4_swap_nodes_in_pairs

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//添加一个前置节点用于统一边界条件的处理
func swapPairs(head *ListNode) *ListNode {
	headPre := &ListNode{
		Next: head,
	}
	for temp := headPre;temp.Next != nil && temp.Next.Next != nil;temp = temp.Next.Next {
		temp.Next, temp.Next.Next, temp.Next.Next.Next = temp.Next.Next, temp.Next.Next.Next, temp.Next
	}
	return headPre.Next
}