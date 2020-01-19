package main

type BinarySearchTree struct {
	Val    int
	Lchild *BinarySearchTree
	Rchild *BinarySearchTree
}

func BuildBinarySearchTree(array []int) *BinarySearchTree {
	var root *BinarySearchTree

	for _, v := range array {
		root = InsertBinarySearchTree(root, v)
	}

	return root
}

func InsertBinarySearchTree(root *BinarySearchTree, val int) *BinarySearchTree {
	if root == nil {
		root = new(BinarySearchTree)
		root.Val = val
		return root
	}

	if val < root.Val {
		root.Lchild = InsertBinarySearchTree(root.Lchild, val)
		return root
	}

	root.Rchild = InsertBinarySearchTree(root.Rchild, val)
	return root
}
