package main

import "fmt"

type BinaryNode struct {
	val   int
	left  *BinaryNode
	right *BinaryNode
}

type List struct {
	node *BinaryNode
	next *List
}

func main() {
	root := &BinaryNode{
		val: 0,
	}

	root = initTree(root)

	list := bfs(root)
	showList(list)
}

func initTree(root *BinaryNode) *BinaryNode {
	head := root

	root.left = &BinaryNode{
		val: 1,
		left: &BinaryNode{
			val: 3,
			left: &BinaryNode{
				val: 7,
			},
			right: &BinaryNode{
				val: 11,
			},
		},
		right: &BinaryNode{
			val: 4,
			left: &BinaryNode{
				val: 8,
			},
		},
	}
	root.right = &BinaryNode{
		val: 2,
		left: &BinaryNode{
			val: 5,
			left: &BinaryNode{
				val: 9,
			},
			right: &BinaryNode{
				val: 12,
			},
		},
		right: &BinaryNode{
			val: 6,
			right: &BinaryNode{
				val: 10,
			},
		},
	}

	return head
}

func bfs(root *BinaryNode) *List {
	result := new(List)
	if root == nil || root.left == nil && root.right == nil {
		result.next = &List{
			node: root,
		}
		return result
	}

	head := result
	var queue []*BinaryNode
	var cur *BinaryNode
	visitM := make(map[*BinaryNode]bool)

	queue = append(queue, root)
	for len(queue) > 0 {
		cur = queue[0]

		if !visitM[cur] {
			visitM[cur] = true

			if cur.left != nil {
				queue = append(queue, cur.left)
			}

			if cur.right != nil {
				queue = append(queue, cur.right)
			}

			l := &List{
				node: cur,
			}
			result.next = l
			result = result.next
		}

		queue = queue[1:]
	}

	return head
}

func showList(l *List) {
	l = l.next

	fmt.Printf("head")
	for l != nil {
		fmt.Printf("->%d", l.node.val)
		l = l.next
	}
	fmt.Println()
}
