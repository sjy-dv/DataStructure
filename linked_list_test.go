package datastructure_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type LinkNode struct {
	value any
	next  *LinkNode
}

type LinkedList struct {
	head   *LinkNode
	tail   *LinkNode
	length int
}

// not use index
// only link to find append anything

// create new node
func NewNode(val any) *LinkNode {
	return &LinkNode{
		value: val,
		next:  nil,
	}
}

// create new linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil, tail: nil, length: 0}
}

// print all node in linkedlist
func (list *LinkedList) printList() {
	if list.length == 0 {
		fmt.Println("list is empty")
	}
	current := list.head
	// GO while is not exists, intend to for loop
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

// append node
func (list *LinkedList) append(val any) {
	node := NewNode(val)
	// is empty, you create node set head, tail
	// becasue only one node! head & tail same
	if list.length == 0 {
		list.head = node
		list.tail = node
		list.length++
		return
	}
	// when already node, you changed tail & last node.next update new node
	list.tail.next = node // is nil => new node
	list.tail = node      // prev tail node => new npde
	list.length++
}

// pop is removed last node
func (list *LinkedList) pop() *LinkNode {
	if list.length == 0 {
		return nil
	}
	// we know last node => is tail node !!
	// but previous node is we dont know

	// using 2 pointer
	pre := list.head
	temp := list.head
	// pre is saved previous node
	// temp is using iterator linked list

	for temp.next != nil {
		pre = temp
		temp = temp.next
	}
	// when find node
	list.tail = pre
	list.tail.next = nil
	list.length--
	// when length is zero
	// becase when 1 node item, for is not working
	// pre, temp is same -> not remove item
	// so, we clear linked_list when length 1=>0
	if list.length == 0 {
		list.head = nil
		list.tail = nil
	}
	return temp
}

// append first node
func (list *LinkedList) prepend(val any) {
	node := NewNode(val)
	//if start! link is zero
	if list.length == 0 {
		list.head = node
		list.tail = node

	} else {
		// we know head
		//step1. head node
		node.next = list.head
		// newnode <- headnode..[other link node]
		list.head = node
		//linked head <- newnode ok!
	}
	list.length++
}

// reverse pop
func (list *LinkedList) popfirst() *LinkNode {
	// when list is zero
	if list.length == 0 {
		return nil
	}
	// we know head node and next node
	temp := list.head
	list.head = temp.next
	// !!! Warning
	// You need to be careful to initialize temp.next.
	//Although it is not used after being returned and
	//will be collected by the GC, retaining the next reference
	//imposes the responsibility of traversing the reference chain
	//on the garbage collector.
	temp.next = nil
	list.length--
	if list.length == 0 {
		list.head = nil
		list.tail = nil
	}
	return temp
}

func (list *LinkedList) get(index int) *LinkNode {
	// filter to invalid index & index range
	if index < 0 || index >= list.length {
		return nil
	}
	temp := list.head
	// jump index step using next
	// node A -> node B -> x index step
	for range index {
		temp = temp.next
	}
	return temp
}

func (list *LinkedList) set(index int, val any) bool {
	temp := list.get(index)
	if temp == nil {
		return false
	}
	temp.value = val
	return true
}

// set is simple update
// but insert is more detail, set index new data
// and next index push one space at a time
func (list *LinkedList) insert(index int, val any) bool {
	if list.length-1 == index {
		list.append(val)
		return true
	} else if index == 0 {
		list.prepend(val)
		return true
	}
	node := NewNode(val)
	temp := list.get(index - 1)
	if temp == nil {
		return false
	}
	node.next = temp.next
	temp.next = node
	list.length++
	return true
}

func (list *LinkedList) remove(index int) *LinkNode {
	if list.length-1 == index {
		return list.pop()
	}
	if index == 0 {
		return list.popfirst()
	}
	temp := list.get(index - 1)
	node := temp.next
	temp.next = node.next
	node.next = nil
	list.length--
	return node
}

func (list *LinkedList) reverse() {
	if list.length == 0 || list.length == 1 {
		return
	}
	// we first changed head & tail
	// if you imagine [1 2 3] Linked List
	temp := list.head
	list.head = list.tail
	list.tail = temp
	// we have
	// temp = 1
	// head = 3
	// tail = 1
	var after *LinkNode = nil
	var before *LinkNode = nil
	for range list.length {
		// after = 1
		// after.next = nil
		after = temp.next
		temp.next = before
		before = temp
		temp = after
		//before 		temp 	after
		//  NIL  		1   	 2   			3

		// ok see.. after = temp.next <- same to view table
		// after 2
		// temp.next => before
		// origin => (1 -> 2)
		// this logic => (1 -> NIL)
		// before = temp is nil replace to temp before(NIL -> 1)
		// temp is 2
		// and temp is 1 => after (after is temp.next <- 2)
		// show changed table
		// before		temp 	after
		//  1    <-      2      3

		// iter 2
		//before  temp 	after
		// 1      2      3
		// after 3
		// temp.next => before
		// origin (2 -> 3) => (2 -> 1)
		// before = temp -> before (1 -> 2)
		// middle table

		//		(before, temp) 		after
		//	1	<-	2				 3
		//temp  = 3

		// iter 3
		// after nil
		// temp.next => before : (3 -> nil) => (3 -> 2)
		// before = temp (2 -> 3)
		// temp  = after (NIL)
		// final view
		// 			 before	 temp	after
		// 1 <-  2  <-   3
	}
}

func TestLinkedListBasic(t *testing.T) {
	linklist := NewLinkedList()
	linklist.append(1)
	linklist.append(2)
	linklist.append(3)
	linklist.printList()

	fmt.Println(linklist.pop())
	linklist.printList()
	fmt.Println(linklist.pop())
	linklist.printList()
	fmt.Println(linklist.pop())
	linklist.printList()
	// more pop list length
	fmt.Println(linklist.pop())
	linklist.printList()
}

func TestLinkedListAdvanced(t *testing.T) {
	linklist := NewLinkedList()
	linklist.append(1)
	linklist.append(2)
	linklist.append(3)
	linklist.printList()

	//
	linklist.prepend("first")
	first := linklist.get(0)
	assert.Equal(t, "first", first.value.(string))

	delfirst := linklist.popfirst()
	assert.Equal(t, first, delfirst)
	first = linklist.get(0)
	assert.Equal(t, 1, first.value.(int))

	ok := linklist.set(0, -1)
	assert.Equal(t, true, ok)
	ok = linklist.insert(1, -1)
	assert.Equal(t, true, ok)
	a1 := linklist.get(0)
	a2 := linklist.get(1)
	assert.Equal(t, a1.value.(int), a2.value.(int))
	d1 := linklist.remove(1)
	assert.Equal(t, a2, d1, "must same (index is same pointer)")
	linklist.printList()
	linklist.reverse()
	linklist.printList()
}
