package main

import (
	"fmt"

	"github.com/carl-leopard/data-structure/utils"
)

type AVLNode struct {
	Val    int
	Lchild *AVLNode
	Rchild *AVLNode

	Height int //节点的高度
}

func main() {
	array := []int{7, 4, 9, 2, 6, 1, 3, 5, 8, 10}

	var avl *AVLNode
	for _, val := range array {
		avl = Insert(avl, val)
		PreOrder(avl)
		fmt.Println()
	}
}

func GetHeight(avl *AVLNode) int {
	if avl == nil {
		return 0
	}

	return utils.Max(GetHeight(avl.Lchild), GetHeight(avl.Rchild)) + 1
}

func Insert(avl *AVLNode, val int) *AVLNode {
	if avl == nil {
		return &AVLNode{
			Val:    val,
			Height: 1,
		}
	}

	if val < avl.Val {
		avl.Lchild = Insert(avl.Lchild, val)
	} else {
		avl.Rchild = Insert(avl.Rchild, val)
	}

	avl.Height = GetHeight(avl)

	return ReBlance(avl)
}

func ReBlance(avl *AVLNode) *AVLNode {
	bf := BalanceFactor(avl)

	//LL
	if bf == 2 && BalanceFactor(avl.Lchild) > 0 {
		return RotateRchild(avl)
	}

	//RR
	if bf == -2 && BalanceFactor(avl.Rchild) < 0 {
		return RotateLchild(avl)
	}

	//LR
	if bf == 2 && BalanceFactor(avl.Lchild) < 0 {
		avl.Lchild = RotateLchild(avl.Lchild)
		return RotateRchild(avl)
	}

	//RL
	if bf == -2 && BalanceFactor(avl.Rchild) > 0 {
		avl.Rchild = RotateRchild(avl.Rchild)
		return RotateLchild(avl)
	}

	return avl
}

func BalanceFactor(avl *AVLNode) int {
	if avl == nil {
		return 0
	}

	return GetHeight(avl.Lchild) - GetHeight(avl.Rchild)
}

//avl.Lchild.GetHeight()-avl.Rchild.GetHeight() = 2
func RotateRchild(avl *AVLNode) *AVLNode {
	Lchild := avl.Lchild

	avl.Lchild = Lchild.Rchild
	Lchild.Rchild = avl

	Lchild.Height = GetHeight(Lchild)
	avl.Height = GetHeight(avl)

	return Lchild
}

//avl.Lchild.GetHeight()-avl.Rchild.GetHeight() = -2
func RotateLchild(avl *AVLNode) *AVLNode {
	Rchild := avl.Rchild

	avl.Rchild = Rchild.Lchild
	Rchild.Lchild = avl

	avl.Height = GetHeight(avl)
	Rchild.Height = GetHeight(Rchild)

	return Rchild
}

func PreOrder(root *AVLNode) {
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
