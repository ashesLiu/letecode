package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{}
	p, d := head, dummy
	for p != nil {
		if p.Val == val {
			d.Next = p.Next
		} else {
			d.Next = p
			d = d.Next
		}
		p = p.Next
	}
	return dummy.Next
}
