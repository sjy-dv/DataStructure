package datastructure_test

import (
	"fmt"
	"testing"
)

type RTreeNode struct {
	value int
	left  *RTreeNode
	right *RTreeNode
}

// recursive tree is recursion self
type RecursiveBinarySearchTree struct {
	root *RTreeNode
}

func NewRTreeNode(val int) *RTreeNode {
	return &RTreeNode{
		value: val,
		left:  nil,
		right: nil,
	}
}

func NewRecursiveBinarySearchTree() *RecursiveBinarySearchTree {
	return &RecursiveBinarySearchTree{root: nil}
}

func (rtree *RecursiveBinarySearchTree) insert(val int) *RTreeNode {
	if rtree.root == nil {
		rtree.root = NewRTreeNode(val)
	}
	return rtree._function_call_r_insert(rtree.root, val)
}

func (rtree *RecursiveBinarySearchTree) _function_call_r_insert(curNode *RTreeNode, value int) *RTreeNode {
	if curNode == nil {
		return NewRTreeNode(value)
	}
	//insert right
	if curNode.value < value {
		curNode.right = rtree._function_call_r_insert(curNode.right, value)
	}
	if curNode.value > value {
		curNode.left = rtree._function_call_r_insert(curNode.left, value)
	}
	// when. insert data is exists return node
	// prevent duplicate (binarysearchtree cant duplicate)
	return curNode
}

func (rtree *RecursiveBinarySearchTree) contains(val int) bool {
	return rtree._function_call_r_contains(rtree.root, val)
}

func (rtree *RecursiveBinarySearchTree) _function_call_r_contains(curNode *RTreeNode, value int) bool {
	if curNode == nil {
		return false
	}
	//full tree
	// A <- value (A) return
	if curNode.value == value {
		return true
	}
	// move rigth
	//  tree features is if value large, append right
	if curNode.value < value {
		return rtree._function_call_r_contains(curNode.right, value)
	} else {
		return rtree._function_call_r_contains(curNode.left, value)
	}
}

func (rtree *RecursiveBinarySearchTree) delete(value int) {
	rtree._function_call_delete(rtree.root, value)
}

func (rtree *RecursiveBinarySearchTree) _function_call_delete(curNode *RTreeNode, value int) *RTreeNode {
	// this condition, value is not exists
	// return nil
	if curNode == nil {
		return nil
	}
	// value smaller than curNode.value
	// move left
	if value < curNode.value {
		//recursion next left node (call stack)
		curNode.left = rtree._function_call_delete(curNode.left, value)
	} else if value > curNode.value {
		// value larger than curnode.value
		curNode.right = rtree._function_call_delete(curNode.right, value)
	} else {
		// imagine this logic is difficult
		// if tree
		/*
				55
				/
			   41  <- find it
			   41 has not leaf(child) and 41 is curNode
			   we return curNode nil
			   tree changed
			   55
			   /
			  nil (delete ok)
		*/
		if curNode.left == nil && curNode.right == nil {
			return nil
		} else if curNode.left == nil {
			/*
							 55
							 / \
				find it	->  41  412
						     \
						      53
							  left is nil, but right has 53
			*/
			curNode = curNode.right // <- chaged node
			/*
				41           53
				 \     ->
				  53
				53 node covered 41. 41 has removed in memory
			*/
		} else if curNode.right == nil {
			curNode = curNode.left
		} else {
			/*
				this case find node has left right full
				now we find minimum value
				because you imagine
									 34
								   /   \
								 30     45
								 /\     / \
					           10 32  40  55
								if delete 34
									 nil (deleted)
									/   \
								 30     45
								 /\     / \
					           10 32  40  55
								this logic (30, 45)
								tree changed
								      45
									/   \
								 30     55
								 /\     /
					           10 32  40
							   if choose left
							    tree changed
								      30
								    /     \
							  32 or 10    45
								 \       / \
					        32 or 10    40  55
							this tree invalid(32 cant left move to right)
							is difficult
							when must left delete tree changed
							       30
								  /  \
								10   45
								     / \
									40  55
								    /
								   32  needs more process
			*/
			minTree := minSelector(curNode.right)
			curNode.value = minTree
			// convered minValue and curNode Next Removed Duplicate original Value Node
			curNode.right = rtree._function_call_delete(curNode.right, minTree)
		}
	}

	return curNode
}

// tree left must small
func minSelector(curNode *RTreeNode) int {
	for curNode.left != nil {
		curNode = curNode.left
	}
	return curNode.value
}

func TestRecursiveBinarySearchTree(t *testing.T) {
	rbst := NewRecursiveBinarySearchTree()
	fmt.Println(rbst.contains(10))
	rbst.insert(2)
	rbst.insert(1)
	rbst.insert(3)
	fmt.Println(rbst.contains(3))
	fmt.Printf(`
	      %v
	     / \
	    %v   %v
	`, valueWrapper(rbst.root), valueWrapper(rbst.root.left), valueWrapper(rbst.root.right))

	rbst.delete(2)

	fmt.Printf(`
	      %v
	     / \
	    %v   %v
	`, valueWrapper(rbst.root), valueWrapper(rbst.root.left), valueWrapper(rbst.root.right))

	// 	 2
	// 	/ \
	// 1   3

	// 	 3
	// 	/ \
	// 1   None
}

func valueWrapper(curNode *RTreeNode) any {
	if curNode != nil {
		return curNode.value
	}
	return "None"
}
