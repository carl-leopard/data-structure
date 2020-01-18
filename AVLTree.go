package main

import (
	"fmt"
	. "github.com/carl-leopard/data-structure/meta"
)

func main() {
	array := []int{1, 5, 9, 10, 100, 2, 5, 8, 16}

}

func (root *AVLNode) insert(val int) *AVLNode {
	if root == nil {
		root = NewAVLNode(val)

		return
	}

	for root != nil {
		if root.Val < val {
			root.Depth -= 1

			if root.RightChild != nil {
				root = root.RightChild
				continue
			}

			break
		}

		if root.Val > val {
			root.Depth += 1

			if root.LeftChild != nil {
				root = root.LeftChild
				continue
			}
			break
		}
	}

	node := NewAVLNode(val)
	node.Parent = root
	if root.Val < val {
		root.RightChild = node
		return node
	}

	root.LeftChild = node
	return node
}

func adjust(root *AVLNode) {
	if root.Depth < 2 && root.Depth > -2 {
		return
	}

	if root.Depth == 2 {
		node := findMaxInLeft(root)

		node.Parent.RightChild = node.LeftChild
		node.RightChild = root
		node.Parent = nil
	}
}

func findMaxInLeft(root *AVLNode) *AVLNode {
	tmp := root.LeftChild
	for tmp != nil && tmp.RightChild != nil {
		tmp = tmp.RightChild
	}

	return tmp
}
