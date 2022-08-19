package linkedlist

var root *ListNode = nil
func reverseList(head *ListNode) *ListNode {
	recursion(head)
	return root
}

func recursion(head *ListNode) *ListNode{
	if head == nil{
		return nil
	}
	root = head
	prev := recursion(head.Next)
	if prev != nil{
		prev.Next = head
	}
	head.Next = nil
	return head
}

func iteration(head *ListNode) *ListNode{
	var root *ListNode = nil
	p := head
	for p!=nil{
		next := p.Next
		p.Next = root
		root = p
		p = next
	}
	return root	
}