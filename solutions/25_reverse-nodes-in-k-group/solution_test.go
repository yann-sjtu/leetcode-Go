package _5_reverse_nodes_in_k_group

import (
	"fmt"
	"testing"
)

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	headPre := &ListNode{
		Next: head,
	}

	temp := headPre
	var flag bool
	for i := 0; i < k; i++ {
		if temp.Next == nil {
			flag = true
			break
		}
		temp = temp.Next
	}
	if flag {
		return headPre.Next
	}
	cur, next := head, head.Next
	for i := 1; i < k; i++ {
		//temp = next.Next
		//next.Next = cur
		//cur = next
		//next = temp
		next.Next, cur, next = cur, next, next.Next
	}
	headPre.Next = cur
	head.Next = reverseKGroup(next, k)
	return headPre.Next
}

func makeList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{}
	temp := head
	for _, n := range nums {
		cur := &ListNode{
			Val: n,
		}
		temp.Next = cur
		temp = temp.Next
	}
	return head.Next
}

func Test_(t *testing.T) {
	list := makeList([]int{1, 2, 3, 4, 5})
	head := reverseKGroup(list, 2)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
