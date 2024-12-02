package datastructure_test

import (
	"fmt"
	"testing"

	"github.com/rs/zerolog/log"
	"golang.org/x/exp/constraints"
)

// i know go using container/heap, more easy to access heap
// but we know how to build heap from scratch

// heap has two types maxheap, minheap

// first maxheap only difference swap

// using generic, not any
// because, any is ok, but the essential tendencies do not match

// when to use heap?
// we find max or min data or sorting
// we can use linked_list but linked list
// search O(n), but heap is O(1) and resorting heap is O(log n)
type MaxHeap[T constraints.Ordered] struct {
	heap []T
}

func NewMaxHeap[T constraints.Ordered]() *MaxHeap[T] {
	return &MaxHeap[T]{
		heap: make([]T, 0),
	}
}

// imagine
// heap looks like binary search tree structure
// but heap first mission is complete tree
// heap always complete tree first, and swap value
// by features of heap (min or max)

/*
    Heap(Max)              BinarySearchTree (The assumption is that this is not possible in the first place, but it can be done.)
      10                       10
    /  \                      /  \
   5    7                   5      7
  /                         /
 9                         9
 <- if add 8

 Chnaged Tree when Add 8
    Heap(max)
    5    7                9   7
   / \      (swap) ->    / \
  9   8                 5   8

  BinarySearchTree

    5    7
   /      \
  9        8
*/

/*
   features : heap always fil tree left -> right

                     10  <- index : 0
                   /    \
                  5 (1)  7  (2) <- heap left -> right fil
                 / \     /\
                9   10  6  4
               (3)  (4) (5) (6)
    rule : left index : y = 2x + 1 (x=index, y=target_left_index)
    right:  y = 2x + 2 // if 10 <- node left right is
    (2x(4)+1, 2x(4)+2) => (9, 10)
    drawing
    9 (7, 8) -> because 4 is index6 and tree complete left to right
    9 has index 7, 8 child 10 has (9, 10). ok cool
*/

func (heap *MaxHeap[T]) _leftChild(index int) int {
	return 2*index + 1
}
func (heap *MaxHeap[T]) _rightChild(index int) int {
	return 2*index + 2
}

func (heap *MaxHeap[T]) _parent(index int) int {
	return (index - 1) / 2
}

// swap (1,2) => (2,1)
// not using heap memory
func (heap *MaxHeap[T]) _swap(a, b int) {
	heap.heap[a], heap.heap[b] = heap.heap[b], heap.heap[a]
}

func (heap *MaxHeap[T]) insert(value T) bool {
	heap.heap = append(heap.heap, value)
	//because append is last node position
	//(if you want to remember detail. reference for linked_list)
	current := len(heap.heap) - 1

	// now heap always sorting to heap types
	// compare parent node
	// and fit conidtion. swap
	// when current == 0 (0 is root. all jobs the end)
	for current > 0 && heap.heap[current] > heap.heap[heap._parent(current)] {
		heap._swap(current, heap._parent(current))
		current = heap._parent(current)
	}
	return true
}

func (heap *MaxHeap[T]) pop() T {
	temp := heap.heap[len(heap.heap)-1] //last node
	heap.heap = heap.heap[:len(heap.heap)-1]
	return temp
}

func (heap *MaxHeap[T]) _sink_down(index int) {
	max := index
	for {
		left := heap._leftChild(max)
		right := heap._rightChild(max)
		// this condition
		// prevent empty index
		// 5
		//3 2 <- this case. 2 left right is empty
		if (left < len(heap.heap)) && heap.heap[left] > heap.heap[max] {
			max = left
		}
		if (right < len(heap.heap)) && heap.heap[right] > heap.heap[max] {
			max = right
		}
		if max != index {
			heap._swap(index, max)
			index = max
		} else {
			return
		}
	}
}

func (heap *MaxHeap[T]) remove() *T {
	if len(heap.heap) == 0 {
		log.Warn().Msg("heap is empty")
		return nil
	}
	if len(heap.heap) == 1 {
		temp := heap.pop()
		return &temp
	}

	temp := heap.heap[0]
	heap.heap[0] = heap.pop()
	heap._sink_down(0)
	return &temp
}

func TestHeap(t *testing.T) {
	heap := NewMaxHeap[int]()

	heap.insert(100)
	heap.insert(241)
	heap.insert(1)
	heap.insert(51)
	heap.insert(42)
	heap.insert(33)
	heap.insert(601)
	heap.insert(214)
	heap.insert(92)

	fmt.Println("max heap status : ", heap.heap)

	a := heap.remove()
	fmt.Println("remove heap node : ", *a)
	fmt.Println("remove after heap status : ", heap.heap)

	// max heap status :  [601 214 241 100 42 1 33 51 92]
	// remove heap node :  601
	// remove after heap status :  [241 214 92 100 42 1 33 51]
}
