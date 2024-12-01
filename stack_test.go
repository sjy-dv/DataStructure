package datastructure_test

import (
	"fmt"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

type StackInNode struct {
	value any
	next  *StackInNode
}

// linkend list is
// [0, 1 ,2 ]

// stack
// LIFO(Last input first out)
// 2 -> top
// 1
// 0

// stack have 2 func (push, pop)
// push continous apped height
// pop remove top
type Stack struct {
	top    *StackInNode
	height int
}

func NewStackInNode(val any) *StackInNode {
	return &StackInNode{
		value: val,
		next:  nil,
	}
}

func NewStack() *Stack {
	return &Stack{top: nil, height: 0}
}

func (s *Stack) stackPrint() {
	if s.height == 0 {
		fmt.Println("stack is empty")
	} else {
		current := s.top
		for current != nil {
			fmt.Println(current.value)
			current = current.next
		}
	}
}

func (s *Stack) push(val any) {
	stack := NewStackInNode(val)
	if s.height == 0 {
		s.top = stack
	} else {
		/*
		   1 <- top
		   0
		    => push 2
		   2 <- top
		   1 <- next
		   0
		*/
		stack.next = s.top
		s.top = stack
	}
	s.height++
}

func (s *Stack) pop() *StackInNode {
	if s.height == 0 {
		log.Warn().Msg("stack height is 0. cant pop")
		return nil
	}
	temp := s.top
	s.top = temp.next
	temp.next = nil
	return temp
}

func TestStack(t *testing.T) {
	stack := NewStack()
	stack.push(0)
	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(4)
	stack.stackPrint()
	sval := stack.pop()
	assert.Equal(t, 4, sval.value.(int))
	stack.stackPrint()
}
