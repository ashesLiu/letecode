package linkedlist

// import "fmt"

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	l, r := head, head.Next
	root, prev := head.Next, l
	for l!=nil{
		if l.Next != nil{
			r = l.Next
			// store next
			next := r.Next
			// switch l and r
			prev.Next = r
			r.Next = l
			l.Next = nil
			prev = l
			// move to next
			l = next
		}else{
			prev.Next = l
			l = l.Next
		}
	}
	return root
}