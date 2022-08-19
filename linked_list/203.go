package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	prev := head
	for prev != nil && prev.Val == val {
		prev = prev.Next
	}
	root := prev
	p := prev.Next
	for p != nil {
		if p.Val == val {
			prev.Next = p.Next
		}
		p = p.Next
	}
	return root
}
