package linkedlist

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	for n>1 {
		fast = fast.Next
		n--
	}
	low := head
	var prev *ListNode
	for fast.Next != nil{
		prev = low
		low = low.Next
		fast = fast.Next
	}
	if prev != nil{
		next := low.next
		prev.Next = next
		low = nil
		return head
	}else{
		return head.Next
	}
}