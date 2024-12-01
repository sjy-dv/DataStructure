package datastructure_test

import "fmt"

type DNode struct {
	value any
	next  *DNode
	prev  *DNode
}

type DoublyLinkedList struct {
	head   *DNode
	tail   *DNode
	length int
}

func NewDNode(val any) *DNode {
	return &DNode{
		value: val,
		next:  nil,
		prev:  nil,
	}
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{head: nil, tail: nil, length: 0}
}

func (list *DoublyLinkedList) printList() {
	if list.length == 0 {
		fmt.Println("list is empty")
	}
	current := list.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}
