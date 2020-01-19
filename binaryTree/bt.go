package main

import (
	"container/list"
	"fmt"
)

//左右子树没有大小之分
type BinaryTree struct {
	Val    int
	Lchild *BinaryTree
	Rchild *BinaryTree
	Parent *BinaryTree
}

func main() {
	array := []int{5, 9, 4, 0, 2, 8, 10}
	root := BuildBinaryTree(array)
	PreOrder(root)
	fmt.Println()
	PreOrderTraverseNoRecursive(root)
	fmt.Println()
	PreOrderTraverseNoQueue(root)
}

func BuildBinaryTree(array []int) *BinaryTree {
	if len(array) == 0 {
		return nil
	}

	bts := make([]*BinaryTree, 0, len(array))
	for _, v := range array {
		node := &BinaryTree{
			Val: v,
		}
		bts = append(bts, node)
	}

	root := bts[0]

	l := list.New()
	l.PushBack(bts[0])

	for i := 1; i < len(bts); {
		e := l.Front()
		node := e.Value.(*BinaryTree)
		l.Remove(e)

		node.Lchild = bts[i]
		bts[i].Parent = node
		l.PushBack(bts[i])
		i++

		node.Rchild = bts[i]
		bts[i].Parent = node
		l.PushBack(bts[i])
		i++
	}

	return root
}

func PreOrder(root *BinaryTree) {
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

func PreOrderTraverseNoRecursive(root *BinaryTree) {
	if root == nil {
		return
	}

	l := list.New()
	var cur, next *BinaryTree
	cur = root
	for cur != nil {
		fmt.Printf("%d ", cur.Val)
		if cur.Rchild != nil {
			l.PushFront(cur.Rchild)
		}

		if cur.Lchild != nil {
			next = cur.Lchild
		} else {
			if l.Len() > 0 {
				e := l.Front()
				l.Remove(e)
				next = e.Value.(*BinaryTree)
			} else {
				next = nil
			}
		}

		cur = next
	}
}

func PreOrderTraverseNoQueue(root *BinaryTree) {
	var pre, cur, next *BinaryTree

	for cur = root; cur != nil; {
		if pre == cur.Parent {
			fmt.Printf("%d ", cur.Val)

			pre = cur
			next = cur.Lchild
		}

		if next == nil || pre == cur.Lchild {
			pre = cur
			next = cur.Rchild
		}

		if next == nil || pre == cur.Rchild {
			pre = cur
			next = cur.Parent
		}

		cur = next
	}
}
