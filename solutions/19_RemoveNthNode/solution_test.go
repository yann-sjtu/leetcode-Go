package _9_RemoveNthNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := &ListNode{
		Next: head,
	}
	lead, behind := temp, temp
	for i := 0; i < n; i++ {
		lead = lead.Next
	}
	for lead.Next != nil {
		lead = lead.Next
		behind = behind.Next
	}
	behind.Next = behind.Next.Next
	return temp.Next
}
