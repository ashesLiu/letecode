package linkedlist

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p, tail := head, head
	var prev *ListNode = nil
	cnt := 0
	for cnt < n {
		tail = tail.Next
		cnt++
	}
	for tail != nil{
		tail = tail.Next
		prev = p
		p = p.Next
	}
	if prev == nil{
		head = p.Next
	}else{
		prev.Next = p.Next
	}
	p = nil
	return head
}