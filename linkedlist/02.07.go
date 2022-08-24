package linkedlist

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    pa, pb := headA, headB
	state := 0
	for pa!=nil&&pb!=nil{
		if pa == pb{
			return pa
		}
		pa = pa.Next
		pb = pb.Next
		if pa == nil{
			pa = headB
			if state == 2{
				break
			}
			state++
		}
		if pb == nil{
			pb = headA
			if state == 2{
				break
			}
			state++
		}
	}
	return nil
}