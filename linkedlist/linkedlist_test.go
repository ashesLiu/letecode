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
