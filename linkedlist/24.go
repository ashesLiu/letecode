package linkedlist

// import "fmt"

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	root := head.Next
	pi, pj := head, head.Next
	var prev *ListNode 
	for pi != nil && pj!=nil {
		pj = pi.Next
		if pj != nil{
			pi,prev = swap(prev, pi, pj)
		}
	}
	return root
}

func swap(prev, pa, pb *ListNode) (*ListNode, *ListNode) {
	next := pb.Next
	if prev != nil {
		prev.Next = pb
	}
	pb.Next = pa
	pa.Next = next
	prev = pa
	return next, prev
}