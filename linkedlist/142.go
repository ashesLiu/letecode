package linkedlist

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil{
		return nil
	}
    l,r := head.Next, head.Next.Next
	for l!=r{
		l = l.Next
		if r.Next == nil || r.Next.Next == nil{
			return nil
		}
		r = r.Next.Next
	}
	r = head
	for l != r{
		l = l.Next
		r = r.Next
	}
	return l
}