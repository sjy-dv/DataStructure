package datastructure_test

import (
	"fmt"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

type QNode struct {
	value any
	next  *QNode
}

type Queue struct {
	first  *QNode
	last   *QNode
	length int
}

// list
// [ 0,  1,  2 ]
// stack [LIFO]
// 2 -> first output
// 1
// 0

// Queue
// last ->  2  1  0 <- first output

func NewQNode(val any) *QNode {
	return &QNode{
		value: val,
	}
}

func NewQueue() *Queue {
	return &Queue{first: nil, last: nil, length: 0}
}

// enqueue like append
// 0 1 <- enequeue(2) : 0 1 2 <- new last
func (q *Queue) enqueue(val any) {
	queue := NewQNode(val)
	if q.length == 0 {
		q.first = queue
		q.last = queue
	} else {
		// new last
		q.last.next = queue
		q.last = queue
	}
	q.length++
}

// dequeue like pop_first()
func (q *Queue) dequeue() *QNode {
	if q.length == 0 {
		log.Warn().Msg("queue length is 0. cant dequeue")
		return nil
	}
	temp := q.first
	if q.length == 1 {
		q.first = nil
		q.last = nil
	} else {
		q.first = q.first.next
		temp.next = nil
	}
	q.length--
	return temp
}

func (q *Queue) printQueue() {
	if q.length == 0 {
		fmt.Println("queue is empty")
	} else {
		current := q.first
		for current != nil {
			fmt.Println(current.value)
			current = current.next
		}
	}
}

func TestQueue(t *testing.T) {
	queue := NewQueue()

	queue.enqueue(0)
	queue.enqueue(1)
	queue.enqueue(2)
	queue.printQueue()
	qval := queue.dequeue()
	assert.Equal(t, 0, qval.value.(int))
	queue.printQueue()
}
