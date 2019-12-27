package __AddTwoNumbers

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	cur := root
	var carry, sum int
	for l1 != nil && l2 != nil {
		sum = l1.Val + l2.Val + carry
		cur.Next = &ListNode{
			Val:  sum % 10,
		}
		carry = sum / 10
		l1 = l1.Next
		l2 = l2.Next
		cur = cur.Next
	}
	for l1 != nil {
		sum = l1.Val + carry
		cur.Next = &ListNode{
			Val:  sum % 10,
		}
		carry = sum / 10
		l1 = l1.Next
		cur = cur.Next
	}
	for l2 != nil {
		sum = l2.Val + carry
		cur.Next = &ListNode{
			Val:  sum % 10,
		}
		carry = sum / 10
		l2 = l2.Next
		cur = cur.Next
	}
	if carry != 0 {
		cur.Next = &ListNode{
			Val:  1,
		}
	}
	return root.Next
}

//padding zero to short list
func addTwoNumbersV2(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	cur := root
	var carry, sum int
	for l1 != nil && l2 != nil {
		sum = l1.Val + l2.Val + carry
		cur.Next = &ListNode{
			Val:  sum % 10,
		}
		carry = sum / 10
		l1 = l1.Next
		l2 = l2.Next
		cur = cur.Next
		if l1 == nil && l2 != nil {
			l1 = &ListNode{}
		}
		if l2 == nil && l1 != nil{
			l2 = &ListNode{}
		}
	}
	if carry != 0 {
		cur.Next = &ListNode{
			Val:  1,
		}
	}
	return root.Next
}