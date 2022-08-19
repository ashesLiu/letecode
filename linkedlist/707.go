package linkedlist

type MyLinkedList struct {
	head, tail *node
	length     int
}

type node struct {
	val  int
	next *node
	prev *node
}

func Constructor() MyLinkedList {
	head := &node{}
	tail := &node{}
	head.next = tail
	tail.prev = head
	return MyLinkedList{head: head, tail: tail}
}

func (this *MyLinkedList) Get(index int) int {
	p := this.findNode(index)
	if p == nil {
		return -1
	} else {
		return p.val
	}
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.insert(val, this.head, this.head.next)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.insert(val, this.tail.prev, this.tail)
}

func (this *MyLinkedList) insert(val int, prev, next *node) {
	newNode := &node{val: val}
	prev.next = newNode
	next.prev = newNode
	newNode.prev = prev
	newNode.next = next
	this.length++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.length {
		return
	}
	if index <= 0 {
		this.AddAtHead(val)
	} else if index == this.length {
		this.AddAtTail(val)
	} else if index < this.length {
		p := this.findNode(index)
		prev := p.prev
		this.insert(val, prev, p)
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	p := this.findNode(index)
	if p != nil {
		this.deleteNode(p)
	}
}

func (this *MyLinkedList) deleteNode(p *node) {
	prev := p.prev
	next := p.next
	prev.next = next
	next.prev = prev
	this.length--
}

func (this *MyLinkedList) findNode(index int) *node {
	if index < 0 || index >= this.length {
		return nil
	}
	p := this.head.next
	cnt := 0
	for cnt < index && p != nil {
		p = p.next
		cnt++
	}
	return p
}
