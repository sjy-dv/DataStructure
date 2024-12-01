package datastructure_test

import (
	"fmt"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

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

func (list *DoublyLinkedList) append(val any) {
	node := NewDNode(val)
	if list.length == 0 {
		list.head = node
		list.tail = node
		list.length++
		return
	}
	list.tail.next = node
	node.prev = list.tail
	list.tail = node
	list.length++
}

func (list *DoublyLinkedList) pop() *DNode {
	if list.length == 0 {
		return nil
	}
	temp := list.tail
	if list.length == 1 {
		list.head = nil
		list.tail = nil
	} else {
		// we can find to easy last node & last one step previous node
		list.tail = list.tail.prev
		list.tail.next = nil
		temp.prev = nil
	}
	list.length--

	return temp
}

func (list *DoublyLinkedList) prepend(val any) {
	node := NewDNode(val)
	if list.length == 0 {
		list.head = node
		list.tail = node
	} else {
		node.next = list.head
		list.head.prev = node
		list.head = node
	}
	list.length++
}

func (list *DoublyLinkedList) popfirst() *DNode {
	if list.length == 0 {
		return nil
	}
	temp := list.head
	list.head = temp.next
	list.head.prev = nil
	temp.next = nil
	list.length--
	if list.length == 0 {
		list.head = nil
		list.tail = nil
	}
	return temp
}

func (list *DoublyLinkedList) get(index int) *DNode {
	if index < 0 || index >= list.length {
		log.Warn().Msg("invalid index")
		return nil
	}
	// we have 2 ways
	// tail -> prev -> find
	// head -> next -> find
	// when? length / 2
	// if 10 => 5 [0~5] [6~10]
	// index lower than middle index -> start head
	// <-> start tail
	middleIndex := int(list.length / 2)
	if middleIndex > index {
		temp := list.head
		for _ = range index {
			temp = temp.next
		}
		return temp
	} else {
		temp := list.tail
		// for condition is different
		// [0 1 2 3 4] // len = 5
		// pos if 3
		// find 2
		// 5(len) - 3(pos)
		for _ = range list.length - 1 - index {
			temp = temp.prev
		}
		return temp
	}
}

func (list *DoublyLinkedList) set(index int, val any) bool {
	if index < 0 || index >= list.length {
		log.Warn().Msg("invalid index")
		return false
	}
	temp := list.get(index)
	if temp == nil {
		return false
	}
	temp.value = val
	return true
}

func (list *DoublyLinkedList) insert(index int, val any) bool {
	if index < 0 || index >= list.length {
		log.Warn().Msg("invalid index")
		return false
	}
	if list.length-1 == index {
		list.append(val)
		return true
	} else if index == 0 {
		list.prepend(val)
		return true
	}
	node := NewDNode(val)
	temp := list.get(index - 1)
	if temp == nil {
		return false
	}
	// 3 node 4
	// node prev -> 3
	// node next -> 4
	node.prev = temp
	node.next = temp.next
	// 4 prev -> node
	temp.next.prev = node
	// 3 next -> node
	temp.next = node
	list.length++
	return true
}

func (list *DoublyLinkedList) remove(index int) *DNode {
	if index < 0 || index >= list.length {
		log.Warn().Msg("invalid index")
		return nil
	}
	if list.length-1 == index {
		return list.pop()
	}
	if index == 0 {
		return list.popfirst()
	}
	temp := list.get(index)
	if temp == nil {
		return nil
	}
	temp.prev.next = temp.next
	temp.next.prev = temp.prev
	temp.next = nil
	temp.prev = nil
	list.length--
	return temp
}

func TestDoublyLinkedListBasic(t *testing.T) {
	dlist := NewDoublyLinkedList()
	dlist.append(0)
	dlist.append(1)
	dlist.append(2)
	dlist.append(3)
	dlist.append(4)
	fmt.Println("======append 0 ~ 4===========")
	dlist.printList()

	fmt.Println("========pop===========")
	fmt.Println(dlist.pop())
	fmt.Println("========pop result===========")
	dlist.printList()
	fmt.Println("========pop first ===========")
	fmt.Println(dlist.popfirst())
	fmt.Println("========pop first result===========")
	dlist.printList()
	fmt.Println("========pop===========")
	fmt.Println(dlist.pop())
	fmt.Println("========pop result===========")
	dlist.printList()
	fmt.Println("========pop first ===========")
	fmt.Println(dlist.popfirst())
	fmt.Println("========pop first result===========")
	dlist.printList()
	fmt.Println("========pop===========")
	fmt.Println(dlist.pop())
	fmt.Println("========pop result===========")
	dlist.printList()
	fmt.Println("========pop first ===========")
	fmt.Println(dlist.popfirst())
	fmt.Println("========pop first result===========")
	dlist.printList()

}

func TestDoublyLinkedListAdvanced(t *testing.T) {
	dlist := NewDoublyLinkedList()
	dlist.append(0)
	dlist.append(1)
	dlist.append(2)
	dlist.append(3)
	fmt.Println("======append 0 ~ 3===========")
	dlist.printList()
	fmt.Println("======prepend 4===========")
	dlist.prepend(4)
	dlist.printList()
	fmt.Println("======case 1===========")
	zero := dlist.get(0)
	mustZero := dlist.popfirst()
	assert.Equal(t, zero.value.(int), mustZero.value.(int))
	fmt.Println("=========case 2==============")
	ok := dlist.set(0, -1)
	assert.Equal(t, true, ok)
	ok = dlist.insert(1, -1)
	assert.Equal(t, true, ok)
	a1 := dlist.get(0)
	a2 := dlist.get(1)
	assert.Equal(t, a1.value.(int), a2.value.(int))
	fmt.Println("length=======>", dlist.length)
	fmt.Println(dlist.remove(3))
	fmt.Println("guess nil with msg", dlist.remove(100))
}
