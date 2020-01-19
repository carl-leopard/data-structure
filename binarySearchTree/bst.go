package main

import "fmt"

type BinarySearchTree struct {
	Val    int
	Lchild *BinarySearchTree
	Rchild *BinarySearchTree
}

func main() {
	array := []int{5, 9, 4, 0, 2, 1, 3, 8, 10}
	root := BuildBinarySearchTree(array)
	PreOrder(root)
	fmt.Println()
	DeleteBst(root, 0)
	PreOrder(root)
	fmt.Println()
	DeleteBst(root, 5)
	PreOrder(root)
	fmt.Println()
}

func BuildBinarySearchTree(array []int) *BinarySearchTree {
	var root *BinarySearchTree

	for _, v := range array {
		root = InsertBinarySearchTree(root, v)
	}

	return root
}

func PreOrder(root *BinarySearchTree) {
	if root == nil {
		return
	}

	fmt.Printf("%d ", root.Val)

	if root.Lchild != nil {
		PreOrder(root.Lchild)
	}

	if root.Rchild != nil {
		PreOrder(root.Rchild)
	}
}

func DeleteBst(root *BinarySearchTree, val int) bool {
	if root == nil {
		return false
	}

	if val == root.Val {
		delete(root)
		return true
	}

	if val < root.Val {
		DeleteBst(root.Lchild, val)
	}

	if val > root.Val {
		DeleteBst(root.Rchild, val)
	}

	return false
}

func delete(node *BinarySearchTree) {
	//删除的节点是叶子节点
	if node.Lchild == nil && node.Rchild == nil {
		node = nil
	}

	//删除的节点只有右孩子
	if node.Lchild == nil && node.Rchild != nil {
		rc := node.Rchild

		node.Val = rc.Val
		node.Lchild = rc.Lchild
		node.Rchild = rc.Rchild
		return
	}

	//删除的节点只有左孩子
	if node.Rchild == nil && node.Lchild != nil {
		lc := node.Lchild

		node.Val = lc.Val
		node.Lchild = lc.Lchild
		node.Rchild = lc.Rchild
		return
	}

	//删除的节点有左右孩子
	pre := node
	rl := node.Lchild
	for rl.Rchild != nil { //找左子树的最大节点
		pre = rl
		rl = rl.Rchild
	}

	node.Val = rl.Val

	if pre != node {
		pre.Rchild = rl.Lchild
	} else {
		pre.Lchild = rl.Lchild
	}
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
