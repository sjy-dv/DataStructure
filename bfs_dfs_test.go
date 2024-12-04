package datastructure_test

import (
	"fmt"
	"testing"

	"github.com/rs/zerolog/log"
)

type TreeNodeCopy struct {
	value int
	left  *TreeNodeCopy
	right *TreeNodeCopy
}

type BinarySearchTreeCopy struct {
	root *TreeNodeCopy
}

func NewTreeNodeCopy(val int) *TreeNodeCopy {
	return &TreeNodeCopy{
		value: val,
		left:  nil,
		right: nil,
	}
}

func NewBinarySearchTreeCopy() *BinarySearchTreeCopy {
	return &BinarySearchTreeCopy{root: nil}
}

func (tree *BinarySearchTreeCopy) insert(val int) bool {
	node := NewTreeNodeCopy(val)
	if tree.root == nil {
		tree.root = node
		return true
	}
	temp := tree.root
	for {
		if temp.value == node.value {
			log.Warn().Msg("tree value must different")
			return false
		}
		// right node case
		if temp.value < node.value {
			if temp.right == nil {
				temp.right = node
				return true
			}
			temp = temp.right
		} else {
			if temp.left == nil {
				temp.left = node
				return true
			}
			temp = temp.left
		}
	}
}

// using datastructure binary search tree
/*
        10
      /   \
     7     30
    /\     /\
   3  8  28  91

   bfs start
   ---1
   queue: [10]
   result []
   ---2
   queue: [10, 7, 30]
   result []
   -> result [10] queue [7, 30]
   ---3
   queue: [7, 30, 3, 8]
   result [10]
   -> result [10, 7] queue [30, 3, 8]
   ---4
   queue: [30, 3, 8, 28, 91]
   result [10, 7]
   -> result [10, 7, 30]
   for queue length until 0
   queue: []
   result [10, 7, 30, 3, 8, 28, 91]
   we looks like tree copy
*/
func (tree *BinarySearchTreeCopy) bfs() []int {
	curNode := tree.root
	queue := make([]*TreeNodeCopy, 0)
	results := make([]int, 0)
	queue = append(queue, curNode)
	for len(queue) > 0 {
		curNode = queue[0]
		queue = queue[1:]
		results = append(results, curNode.value)

		if curNode.left != nil {
			queue = append(queue, curNode.left)
		}
		if curNode.right != nil {
			queue = append(queue, curNode.right)
		}
	}
	return results
}

// dfs preOrder is find left tree first after right
func (tree *BinarySearchTreeCopy) dfsPreOrder() []int {
	results := make([]int, 0)

	var traverse func(curNode *TreeNodeCopy)
	traverse = func(curNode *TreeNodeCopy) {
		//
		/*
					        10
					      /   \
					     7     30
					    /\     /\
					   3  8  28  91
			           DFS-PreOrder Start
			           result = [10]
			            --traverse start
			            result = [10, 7] - first line travel
			            result = [10, 7, 3] - firstline travel in travel(firstline)
			            result= [10,7,3,8] = firstline travel in travel(seconedline)
			            result = [10,7,3,8,30] - second line travel
			            result = [...30, 28] - second line travel in travel(firstline)
			            result = [...30, 28, 91] - second line travel in travel(secondline)

			            not satisfy condition traverse exit
		*/
		results = append(results, curNode.value)
		if curNode.left != nil {
			traverse(curNode.left)
		}
		if curNode.right != nil {
			traverse(curNode.right)
		}
	}
	traverse(tree.root)
	return results
}

func (tree *BinarySearchTreeCopy) dfsPostOrder() []int {
	results := make([]int, 0)

	var traverse func(curNode *TreeNodeCopy)
	traverse = func(curNode *TreeNodeCopy) {
		//
		/*
					        10
					      /   \
					     7     30
					    /\     /\
					   3  8  28  91
			           DFS-PostOrder Start
			           --traverse start
			            result = [3] - first line travel
			            result = [3, 8, 7] - firstline travel in travel(firstline)
			            result = [3, 8, 7, 28, 91, 30] - second line travel
			            result = [... 10] - last root node

			            not satisfy condition traverse exit
		*/
		if curNode.left != nil {
			traverse(curNode.left)
		}
		if curNode.right != nil {
			traverse(curNode.right)
		}
		results = append(results, curNode.value)
	}
	traverse(tree.root)
	return results
}

func (tree *BinarySearchTreeCopy) dfsInOrder() []int {
	results := make([]int, 0)

	var traverse func(curNode *TreeNodeCopy)
	traverse = func(curNode *TreeNodeCopy) {

		if curNode.left != nil {
			traverse(curNode.left)
		}
		/*
					        10
					      /   \
					     7     30
					    /\     /\
					   3  8  28  91
			           DFS-PostOrder Start
			           --traverse start
			            result = [3] - first line travel
			            result = [3, 7, 8] - firstline travel in travel(firstline)
					     //append
					     result = [3, 7, 8, 10]
			            result = [3, 8, 7,, 10 28, 30, 91] - second line travel
			            not satisfy condition traverse exit
		*/
		results = append(results, curNode.value)
		if curNode.right != nil {
			traverse(curNode.right)
		}
	}
	traverse(tree.root)
	return results
}

func TestBfs(t *testing.T) {
	tree := NewBinarySearchTreeCopy()

	tree.insert(10)
	tree.insert(8)
	tree.insert(6)
	tree.insert(9)
	tree.insert(14)
	tree.insert(13)
	tree.insert(19)
	fmt.Println(tree.bfs())
}

func TestDfsPreOrder(t *testing.T) {
	tree := NewBinarySearchTreeCopy()

	tree.insert(10)
	tree.insert(8)
	tree.insert(6)
	tree.insert(9)
	tree.insert(14)
	tree.insert(13)
	tree.insert(19)
	fmt.Println(tree.dfsPreOrder())
}

func TestDfsPostOrder(t *testing.T) {
	tree := NewBinarySearchTreeCopy()

	tree.insert(10)
	tree.insert(8)
	tree.insert(6)
	tree.insert(9)
	tree.insert(14)
	tree.insert(13)
	tree.insert(19)
	fmt.Println(tree.dfsPostOrder())
}

func TestDfsInOrder(t *testing.T) {
	tree := NewBinarySearchTreeCopy()

	tree.insert(10)
	tree.insert(8)
	tree.insert(6)
	tree.insert(9)
	tree.insert(14)
	tree.insert(13)
	tree.insert(19)
	fmt.Println(tree.dfsInOrder())
}
