package datastructure_test

import (
	"fmt"
	"testing"

	"github.com/rs/zerolog/log"
)

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

type BinarySearchTree struct {
	root *TreeNode
}

// list also tree but only straight tree

//we learned this code. -> binary tree
/*
   binary tree structure

            NodeA            <- parent
           /     \
        NodeB   NodeC      <- child
       /
   NodeD                   <- leaf

   tree have 3 status => full, perfect, complete

   case 1.
       Status
           - Complete All Tree is Full
           - Full Tree is Full
           - Perfect AllTree is Full
       A
      / \
    B   C

   case2.
           Status
               - Complete
               - Full
               - Perfect(X) -> C Tree is not full. Not Perfect
       A
      /  \
     B    C
    /\
   D E
*/

// Compare LinkedList Vs Trees
/*
   Data: [0 1 2 3 4 5]
   FindConditional: 3

   LinkedList 0(head) next-> 1 next -> 2 next -> 3
   Complexity : O(n)

   Trees
        2
       / \
      0   4
     /    /\
    1    3  5

    ▼
     4
    /
   3 <- find it
   Complexity : O(log n)

   applicable job is [search, remove]

   but insert is linkedlist more fast
   linkedlist have continuous characteristics
   only append last [Complexity O(1)]
*/

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		value: val,
		left:  nil,
		right: nil,
	}
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{root: nil}
}

// tree insert feature
// when new node bigger than parent node
// insert right, but smaller => insert left
func (tree *BinarySearchTree) insert(val int) bool {
	node := NewTreeNode(val)
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

func (tree *BinarySearchTree) contains(val int) bool {
	if tree.root == nil {
		log.Warn().Msg("Tree is Empty")
	}
	temp := tree.root
	for temp != nil {
		if temp.value < val {
			temp = temp.right
		} else if temp.value > val {
			temp = temp.left
		} else {
			// temp.value == val
			return true
		}
	}
	// tree not cotains user input
	return false
}

func TestBinarySearchTree(t *testing.T) {
	bst := NewBinarySearchTree()
	fmt.Println(bst.contains(10))
	randSeeds := []int{10, 5, 6, 13, 4, 74, 315, 22, 72}
	for _, seed := range randSeeds {
		fmt.Println(bst.insert(seed))
	}
	fmt.Println(bst.contains(214))
	fmt.Println(bst.contains(74))
}
