package linkedlist

import (
	"fmt"
	"testing"
)

func Test707(t *testing.T) {
	mll := Constructor()
	mll.AddAtIndex(0, 1)
	mll.AddAtTail(2)
	fmt.Printf("mll.findNode(2).val: %v\n", mll.findNode(0).val)
}

func Test24(t *testing.T){
	root := prepareList()
	root = swapPairs(root)
	fmt.Println("output")
	printListNode(root)
}

func printListNode(root *ListNode){
	cnt := 0
	stackSize := 1000
	for root != nil{
		fmt.Printf("%v\t",root.Val)
		cnt++
		if cnt>stackSize{
			break
		}
		root = root.Next 
	}
	fmt.Println()
}

func prepareList() *ListNode{
	root := addNode(nil, 4)
	root = addNode(root, 3)
	root = addNode(root, 2)
	root = addNode(root, 1)
	root = addNode(root, 0)
	return root
}

func addNode(root *ListNode, val int) *ListNode{
	p := &ListNode{}
	p.Val = val
	p.Next = root
	return p
}