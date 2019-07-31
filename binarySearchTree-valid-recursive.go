package main

//valid binary search tree

type Node struct {
	val int
	left *Node
	right *Node
}

func validBinarySearchTree(root *Node, min int, max int) bool {
	if root == nil {
		return true
	}

	if root.val < min || root.val > max {
		return false
	}

	if root.left != nil && root.left.val > root.val {
		return false
	}

	if root.right != nil && root.right.val < root.val {
		return false
	}

	return validBinarySearchTree(root.left, min, root.val) && validBinarySearchTree(root.right, root.val, max)
}